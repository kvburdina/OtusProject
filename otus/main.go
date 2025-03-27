package main

import (
	"fmt"
	"strings"
)

func main() {

	n := 8
	black := "#"
	white := " "
	current := black
	iswhite := false

	builder := strings.Builder{}

	for i := 1; i <= n; i++ {

		for j := 1; j <= n; j++ {

			builder.WriteString(current)

			if n%2 == 0 && j != n || n%2 != 0 {
				iswhite = !iswhite
			}
			if iswhite {
				current = white
			} else {
				current = black
			}
		}
		builder.WriteString("\n")
	}
	fmt.Println(builder.String())
}
