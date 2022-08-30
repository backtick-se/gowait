package task_test

import (
	"context"
	"time"

	. "github.com/backtick-se/gowait/core/task"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Task Queue", func() {
	var err error
	var instance Instance
	var queue TaskQueue
	var image = "img"

	BeforeEach(func() {
		queue = NewTaskQueue(1)
		instance = NewInstance(&Spec{
			Image: image,
		})
	})

	It("queues instances based on image name", func() {
		err = queue.Queue(context.Background(), instance)
		Expect(err).ToNot(HaveOccurred())

		// use a cancelled context so that aquire wont wait
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		// aquire another image should fail
		_, err = queue.Aquire(ctx, "another image")
		Expect(err).To(MatchError(context.Canceled))

		// first aquire should succeed
		aquired, err := queue.Aquire(context.Background(), image)
		Expect(err).ToNot(HaveOccurred())
		Expect(aquired.ID()).To(Equal(instance.ID()))

		// second aquire should fail
		_, err = queue.Aquire(ctx, image)
		Expect(err).To(MatchError(context.Canceled))
	})

	It("aquires as soon as an instance is available", func() {
		go func() {
			time.Sleep(10 * time.Millisecond)
			err = queue.Queue(context.Background(), instance)
			Expect(err).ToNot(HaveOccurred())
		}()

		// aquire should eventually succeed
		aquired, err := queue.Aquire(context.Background(), image)
		Expect(err).ToNot(HaveOccurred())
		Expect(aquired.ID()).To(Equal(instance.ID()))
	})

	It("cancels aquire on context timeout", func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()

		// should fail after 1 ms
		_, err := queue.Aquire(ctx, image)
		Expect(err).To(MatchError(context.Canceled))
	})
})
