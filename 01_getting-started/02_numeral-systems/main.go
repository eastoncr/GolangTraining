package main

import "fmt"

func main() {
	fmt.Println(42)    // print decimal
	fmt.Println(100.2) // print decimal
	fmt.Printf("%d - %b\n", 42, 42)

	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
}
