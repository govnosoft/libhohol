package main

import (
	"bufio"
	"os"
	"unicode"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	i := 0
	for {
		c, _, _ := r.ReadRune()
		if unicode.IsLetter(c) {
			switch i {
			case 0:
				w.WriteRune('h')
				i++
			case 1:
				w.WriteRune('r')
				i++
			default:
				w.WriteRune('u')
			}
		} else {
			w.WriteRune(c)
			i = 0
		}
	}
}
