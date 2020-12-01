package main

import (
	"fmt"
)

func main() {
	integer := 0
	var name string
	fmt.Println("\n----------\nhehehhee\n----------")
	for i := 0; i < 8; i++ {
		fmt.Println("x:", integer)
		increment(&integer)
	}
	fmt.Println("ok gimme ur first name i promise i wont steal ur identity")
	fmt.Scanln(&name)
	fmt.Println("HAHA, ", name, " I HAVE STOLEN UR IDENTITITYTYI!")
}

func increment(x *int) {
	*x++
}
