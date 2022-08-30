package spy_test

import (
	. "github.com/backtick-se/gowait/util/spy"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"fmt"
	"testing"
)

func TestSpy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Util: Spy")
}

var _ = Describe("Spies", func() {
	Context("returns void", func() {
		It("records calls", func() {
			spy := Void{}
			spy.Call(1, "hej")
			spy.Call(2, "hej")

			Expect(spy.Called()).To(BeTrue())
			Expect(spy.CalledWith(1, "hej")).To(BeTrue())
			Expect(spy.Calls()).To(Equal(2))
		})
	})

	Context("returns error", func() {
		It("records calls", func() {
			spy := E{}
			err := spy.Call(1, "hej")
			Expect(spy.Called()).To(BeTrue())
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("returns error if set", func() {
			spy := E{Error: fmt.Errorf("shit")}
			err := spy.Call("hej")
			Expect(spy.Called()).To(BeTrue())
			Expect(err).Should(HaveOccurred())
		})
	})
})
