package greeting

import (
	"bufio"
	"fmt"
	"os"
)

func Scanner() {
	fmt.Println("Use --help to see all commands ")
	for {
		var x int64
		var y int64
		fmt.Println("Use exit to exit calculator")
		fmt.Print("Pick ur chose: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		chose := scanner.Text()

		switch chose {
		case "--help":
			fmt.Println(" + : adds the two numbers together")
			fmt.Println(" - : subtracts the second number from the first")
			fmt.Println(" * : multiplies the first number by the second")
			fmt.Println(" / : divides the first number by the second")
		case "-":
			fmt.Print("Chose first number: ")
			fmt.Scanln(&x)
			fmt.Print("Chose second number: ")
			fmt.Scanln(&y)
			c := x - y
			fmt.Println("Result: ", c)
		case "+":
			fmt.Print("Chose first number: ")
			fmt.Scanln(&x)
			fmt.Print("Chose second number: ")
			fmt.Scanln(&y)
			c := x + y
			fmt.Println("Result: ", c)
		case "*":
			fmt.Print("Chose first number: ")
			fmt.Scanln(&x)
			fmt.Print("Chose second number: ")
			fmt.Scanln(&y)
			c := x * y
			fmt.Println("Result: ", c)
		case "/":
			fmt.Print("Chose first number: ")
			fmt.Scanln(&x)
			fmt.Print("Chose second number: ")
			fmt.Scanln(&y)
			c := x / y
			fmt.Println("Result: ", c)
		case "exit":
			fmt.Println("Goodbye")
			return
		default:
			fmt.Println("ERROR")
		}
	}
}
