package daemon_test

import (
	"context"
	"time"

	"github.com/backtick-se/gowait/core/daemon"
	. "github.com/backtick-se/gowait/core/daemon"
	"github.com/backtick-se/gowait/util/spy"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/backtick-se/gowait/core/cluster"
	"github.com/backtick-se/gowait/core/executor"
	"github.com/backtick-se/gowait/core/task"
)

var _ = Describe("Worker", func() {

	var worker Worker
	var driver *cluster.DriverMock

	BeforeEach(func() {
		driver = &cluster.DriverMock{}
		worker = NewWorker(driver, task.GenerateID("test"), "image")
	})

	Context("state transitions", func() {
		BeforeEach(func() {
			err := worker.Start(context.Background())
			Expect(err).ToNot(HaveOccurred())
			Expect(worker.Status()).To(Equal(executor.StatusWait))

			Expect(driver.SpawnSpy).To(spy.Called())

			err = worker.OnInit()
			Expect(err).ToNot(HaveOccurred())
			Expect(worker.Status()).To(Equal(executor.StatusIdle))
		})

		It("rejects init once initialized", func() {
			err := worker.OnInit()
			Expect(err).To(MatchError(daemon.ErrInvalidState))
		})

		It("aquires tasks", func() {
			// pass a task instance to the worker
			instance := task.NewInstance(&task.Spec{Name: "hello", Image: "image"})
			err := worker.OnAquire(instance)
			Expect(err).ToNot(HaveOccurred())

			// trigger task completion
			err = instance.OnComplete(&task.MsgComplete{})
			Expect(err).ToNot(HaveOccurred())

			// give it a chance to react
			time.Sleep(time.Millisecond)

			// worker should be idle
			Expect(worker.Status()).To(Equal(executor.StatusIdle))
		})

		AfterEach(func() {
			worker.OnStop()
			Expect(worker.Status()).To(Equal(executor.StatusStop))
		})
	})
})
