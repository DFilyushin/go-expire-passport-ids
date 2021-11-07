package fget

import (
	"fmt"
	"io"
)

type ProgressReader struct {
	total int64
	io.Reader
	processed int64
}

func (pt *ProgressReader) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)
	pt.processed += int64(n)

	if err == nil {
		if pt.processed == pt.total {
			fmt.Printf("\rOn 100/100")
		} else {
			v := pt.processed * 100 / pt.total
			fmt.Printf("\rOn %d/100", v)
		}
	}
	return n, err
}
