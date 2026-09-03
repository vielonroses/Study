package main

import "fmt"

func main() {
	var name, fam string
	fmt.Print("Hey, enter ur name: ")
	fmt.Scanln(&name)
	fmt.Println("Heey", EnterName(name, ""), "how are u?")
	fmt.Print("Enter ur name: ")
	fmt.Scanln(&fam)
	fmt.Println("Heey", EnterName(name, fam), "how are u?")
}

func EnterName(name, fam string) string {

	return name + " " + fam
}
