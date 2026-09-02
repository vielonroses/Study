package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	fmt.Println("Parodia on linux 'random' haha ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("sh 5.0 $: ")
		scanner.Scan()
		cmd := scanner.Text()
		switch cmd {
		case "help":
			fmt.Println("Command:")
			fmt.Println("-	random")
			fmt.Println("-	exit")
		case "random":
			fmt.Println(rand.Intn(100000))
		case "exit":
			return
		default:
			if cmd != "" {
				fmt.Println("command not found, please try again")
			}
		}
	}
}
