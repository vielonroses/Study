package main

import (
	"fmt"
	"time"
)

func main() {
	var UserChoose string
	fmt.Print("Wanna know time rn? ")
	fmt.Scanln(&UserChoose)
	if UserChoose == "yes" {
		Oclock()
	} else {
		fmt.Println("Goodbye")
	}
}

func Oclock() {
	for {
		fmt.Println("Time right now: ", time.Now())
		time.Sleep(time.Second)
	}
}
