package executor_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"cowait/executor"
)

var _ = Describe("Process", func() {
	It("does stuff", func() {
		proc, err := executor.Exec("echo", "hello", "world")
		Expect(err).ToNot(HaveOccurred())

		proc.Stdout.Read()

		proc.Wait()
	})
})
