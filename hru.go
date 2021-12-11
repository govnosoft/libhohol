package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	b := strings.Builder{}
	for {
		message, _ := reader.ReadString('\n')
		b.Grow(len(message))
		i := 0
		for _, x := range message {
			if unicode.IsLetter(x) {
				switch i {
				case 0:
					b.WriteRune('h')
					i++
				case 1:
					b.WriteRune('r')
					i++
				default:
					b.WriteRune('u')
				}
			} else {
				b.WriteRune(x)
				i = 0
			}
		}
		fmt.Print(b.String())
		b.Reset()

	}
}
