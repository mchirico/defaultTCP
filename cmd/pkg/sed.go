package pkg

import (
	"fmt"
)
import "io"

func ReadSTDIN(input *io.PipeReader) {

	buf := make([]byte, 100024)

	n, err := io.ReadAtLeast(input, buf, 3)
	fmt.Printf("n= %v %v\n",n, err)

}

