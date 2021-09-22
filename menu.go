package main

import "fmt"

func showMenu() {
	fmt.Println("Menu of the Restaurant is:")
	fmt.Println()

	menu := map[string]float64{
		"apple pie": 3.99,
		"donuts":    0.99,
		"cake":      6.99,
		"muffins":   0.79,
		"pasta":     4.99,
		"Veg Salad": 5.99,
		"Brownie":   2.99,
	}

	// fmt.Println(menu)
	// fmt.Println(menu["donuts"])

	//looping maps

	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

}
