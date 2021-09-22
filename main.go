package main

import "fmt"

//Bill generator for a Restaurant
func main() {

	showMenu()
	fmt.Println()
	generateBill := newBill("Seema's Bill")
	fmt.Println(generateBill)

}
