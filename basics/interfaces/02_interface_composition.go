package interfaces

import "fmt"

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type StringBuffer struct {
	data []byte
}

func (sb *StringBuffer) Read(p []byte) (int, error) {
	n := copy(p, sb.data)
	sb.data = sb.data[n:]
	return n, nil
}

func (sb *StringBuffer) Write(p []byte) (int, error) {
	sb.data = append(sb.data, p...)
	return len(p), nil
}

func ExampleComposition() {
	var rw ReadWriter = &StringBuffer{}

	rw.Write([]byte("hello"))
	buf := make([]byte, 5)
	rw.Read(buf)

	fmt.Println(string(buf))
}
