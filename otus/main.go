package main

import (
	"fmt"
	"strings"
)

const (
	black = "#"
	white = " "
)

func main() {

	var n int
	fmt.Print("Введите размер шахматной доски: ")
	fmt.Scanf("%d", &n)

	builder := strings.Builder{}

	for i := 1; i <= n; i++ {

		for j := 1; j <= n; j++ {

			if (i+j)%2 == 0 {
				builder.WriteString(black)
			} else {
				builder.WriteString(white)
			}
		}

		builder.WriteString("\n")
	}

	fmt.Println(builder.String())
}
