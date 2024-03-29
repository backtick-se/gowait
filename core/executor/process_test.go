package executor_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/backtick-se/gowait/core/executor"
)

var _ = Describe("Process", func() {
	It("does stuff", func() {
		proc, err := executor.Exec("echo", "hello", "world")
		Expect(err).ToNot(HaveOccurred())

		proc.Stdout().Read()

		<-proc.Done()
	})
})
