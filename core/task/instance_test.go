package task_test

import (
	"fmt"
	"time"

	. "github.com/backtick-se/gowait/core/task"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Task Instance", func() {
	var spec *Spec
	var instance Instance
	var done chan struct{}

	BeforeEach(func() {
		spec = &Spec{
			Name:  "cowait.builtin.enumerate",
			Image: "cowait/cowait",
		}
		instance = NewInstance(spec)
	})

	It("starts in the waiting state", func() {
		Expect(instance.State().Status).To(Equal(StatusWait))
	})

	Context("state transitions", func() {
		BeforeEach(func() {
			done = make(chan struct{})
			instance.Start(done)

			err := instance.OnInit(&MsgInit{
				Header:   Header{ID: string(instance.ID()), Time: time.Now()},
				Version:  "1",
				Executor: "asd",
			})
			Expect(err).ToNot(HaveOccurred())

			Expect(instance.State().Status).To(Equal(StatusExec))
			Expect(instance.State().Executor).To(Equal(ID("asd")))
		})

		AfterEach(func() {
			select {
			case <-done:
				break
			default:
				Fail("expected instance routine exit")
			}
		})

		It("transitions to the completed state", func() {
			result := Result("{}")
			err := instance.OnComplete(&MsgComplete{
				Header: Header{ID: string(instance.ID()), Time: time.Now()},
				Result: result,
			})
			Expect(err).ToNot(HaveOccurred())

			Expect(instance.State().Status).To(Equal(StatusDone))
			Expect(instance.State().Result).To(Equal(result))
		})

		It("transitions to the failed state", func() {
			taskerr := fmt.Errorf("shit")
			err := instance.OnFailure(&MsgFailure{
				Header: Header{ID: string(instance.ID()), Time: time.Now()},
				Error:  fmt.Errorf("shit"),
			})
			Expect(err).ToNot(HaveOccurred())

			Expect(instance.State().Status).To(Equal(StatusFail))
			Expect(instance.State().Err).To(Equal(taskerr))
		})
	})
})
