package main

import (
	"bufio"
	"os"
	"unicode"
)

const BUFSIZE = 1048576

func main() {
	r := bufio.NewReaderSize(os.Stdin, BUFSIZE)
	w := bufio.NewWriterSize(os.Stdout, BUFSIZE)
	var i uint8
	for {
		c, _, _ := r.ReadRune()
		if unicode.IsLetter(c) {
			switch i {
			case 0:
				w.WriteRune('h')
				i = 1
			case 1:
				w.WriteRune('r')
				i = 2
			default:
				w.WriteRune('u')
			}
		} else {
			w.WriteRune(c)
			i = 0
		}
	}
}
