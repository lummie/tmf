package ioext

import (
	"io"
	"bytes"
)

const CountBuffer = 32*1024

func CountLines(r io.Reader, separator []byte) (int, error) {
	buf := make([]byte, CountBuffer)
	count := 0

	for {
		c, err := r.Read(buf)
		if c == 0 && err == io.EOF{ // if empty source return zero
			return 0,nil
		}
		count += bytes.Count(buf[:c], separator)

		switch {
		case err == io.EOF:
			return count+1, nil

		case err != nil:
			return count, err
		}
	}
}
