package main

import (
	"fmt"
	"io"
	"os"
)

// Capper implements io.Writer and turns everything to uppercase.
type Capper struct {
	w io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')

	out := make([]byte, len(p))
	for i, b := range p {
		if b >= 'a' && b <= 'z' {
			b -= diff
		}
		out[i] = b
	}
	return c.w.Write(out)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintf(c, "hello, world\n")
}
