package main

import (
	"bufio"
	"os"
	"unicode"

)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	for {
		message, _ := reader.ReadString('\n')
		i := 0
		for _, x := range message {
			if unicode.IsLetter(x) {
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
				w.WriteRune(x)
				i = 0
			}
		}
	}
}
