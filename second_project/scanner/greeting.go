package greeting

import (
	"bufio"
	"fmt"
	"os"
)

func Greeting() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter ur name and 2-nd name: ")
	scanner.Scan()
	input := scanner.Text()

	if input != "" {
		fmt.Printf("Yooo %s, how are u?\n", input)
	} else {
		fmt.Println("Error")
	}
}
