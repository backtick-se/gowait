package executor_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"cowait/core/executor"
	"io"
)

var _ = Describe("Line Pump", func() {
	It("correctly buffers lines", func() {
		testread := NewTestReader(
			[]byte("\n\n"),
			[]byte("hello\nw"),
			[]byte("orld"),
		)
		pump := executor.NewLineReader(testread)

		Expect(pump.Read()).To(Equal("\n"))
		Expect(pump.Read()).To(Equal("\n"))
		Expect(pump.Read()).To(Equal("hello\n"))
		Expect(pump.Read()).To(Equal("world"))

		_, err := pump.Read()
		Expect(err).To(MatchError(io.EOF))

		Expect(pump.Wait()).ToNot(HaveOccurred())
	})
})

type testreader struct {
	reads [][]byte
	pos   int
}

func NewTestReader(reads ...[]byte) io.ReadCloser {
	return &testreader{
		reads: reads,
		pos:   0,
	}
}

func (t *testreader) Read(buf []byte) (int, error) {
	if t.pos >= len(t.reads) {
		return 0, io.EOF
	}
	copy(buf, t.reads[t.pos])
	t.pos += 1
	return len(t.reads[t.pos-1]), nil
}

func (t *testreader) Close() error {
	return nil
}
