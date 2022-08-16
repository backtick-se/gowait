package client_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"cowait/client"
	"io"
)

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

var _ = Describe("Reader Pump", func() {
	It("correctly buffers lines", func() {
		testread := NewTestReader(
			[]byte("hello\nwo"),
			[]byte("rld"),
			[]byte("\n\n"),
		)
		output := make(chan string)
		go client.ReaderPump(testread, output)

		Expect(<-output).To(Equal("hello\n"))
		Expect(<-output).To(Equal("world\n"))
		Expect(<-output).To(Equal("\n"))
		Expect(<-output).To(Equal("\n"))
		_, ok := <-output
		Expect(ok).To(BeFalse())
		Expect(output).To(BeClosed())
	})
})
