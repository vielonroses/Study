package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
)

func main() {
	fmt.Println("Parodia on linux 'random' haha ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("")
		fmt.Print("sh 5.0 $: ")
		scanner.Scan()
		cmd := scanner.Text()
		switch cmd {
		case "help":
			fmt.Println("Command:")
			fmt.Println("-	random")
			fmt.Println("-	exit")
			fmt.Println("-	Pi")
			fmt.Println("-	square")
		case "random":
			fmt.Println(rand.Intn(100000))
		case "exit":
			return
		case "Pi":
			fmt.Printf("Pi: %g", math.Pi)
		case "square":
			var x float64
			fmt.Print("Enter ur number: ")
			fmt.Scanln(&x)
			fmt.Printf("Square root: %g", math.Sqrt(x))

		default:
			if cmd != "" {
				fmt.Println("command not found, please try again.")
			}
		}
	}
}
