package spy_test

import (
	. "github.com/backtick-se/gowait-cloud/util/spy"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Matchers", func() {
	var spy *Void

	BeforeEach(func() {
		spy = &Void{}
	})

	Context("called with", func() {
		It("matches with same arguments", func() {
			spy.Call(1, "hej")
			Expect(spy).To(Called())
		})

		It("rejects wrong arguments", func() {
			Expect(spy).ToNot(Called())
		})

		It("works on E1", func() {
			e1 := E1[int]{}
			Expect(e1).ToNot(Called())
		})
	})

	Context("called with", func() {
		It("matches with same arguments", func() {
			spy.Call(1, "hej")
			Expect(spy).To(CalledWith(1, "hej"))
		})

		It("rejects wrong arguments", func() {
			spy.Call(1, "hej")
			Expect(spy).ToNot(CalledWith(1))
		})
	})
})
