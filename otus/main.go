package main

import (
	"fmt"
	"strings"
)

func main() {

	n := 8
	black := "#"
	withe := " "
	current := black
	isWithe := false

	builder := strings.Builder{}

	for i := 1; i <= n; i++ {

		for j := 1; j <= n; j++ {

			builder.WriteString(current)

			if n%2 == 0 && j != n || n%2 != 0 {
				isWithe = !isWithe
			}
			if isWithe {
				current = withe
			} else {
				current = black
			}
		}
		builder.WriteString("\n")
	}
	fmt.Println(builder.String())
}
