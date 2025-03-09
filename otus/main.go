package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Ksu"
	sname := "Bur"
	fmt.Printf("Hello everybody. My name is %s %s.\n", name, sname)
	currentTime := time.Now()
	weekday := currentTime.Weekday()
	fmt.Printf("Today is %s", weekday)
}
