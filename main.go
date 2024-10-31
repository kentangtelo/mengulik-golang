package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Hello World")
	}

	var i int = 0

	for i<5 {
		fmt.Println("Angka", i)
		i++
	}

	i=0

	for{
		fmt.Println("Jelek", i)

		i++
		if i == 5 {
			break
		}
	}

}
