package pkg

import (
	. "io"
	"testing"
)

// ref: https://golang.org/src/io/pipe_test.go
// ref: https://www.zupzup.org/io-pipe-go/
// ref: https://gist.github.com/ifels/10392762
func checkWrite(t *testing.T, w Writer, data []byte, c chan int) {
	n, err := w.Write(data)
	if err != nil {
		t.Errorf("write: %v", err)
	}
	if n != len(data) {
		t.Errorf("short write: %d != %d", n, len(data))
	}
	c <- 0
}

func Test_Sed(t *testing.T) {
	c := make(chan int)
	r, w := Pipe()
	var buf = make([]byte, 64)
	go checkWrite(t, w, []byte("hello, world"), c)
	n, err := r.Read(buf)
	if err != nil {
		t.Errorf("read: %v", err)
	} else if n != 12 || string(buf[0:12]) != "hello, world" {
		t.Errorf("bad read: got %q", buf[0:n])
	}
	<-c
	r.Close()
	w.Close()

}
