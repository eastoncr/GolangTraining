package main

import "fmt"

func main() {
	var small int
	var big int
	fmt.Println("Please enter a small number followed by a larger number")
	fmt.Scan(&small, &big)
	println(big % small)
}
