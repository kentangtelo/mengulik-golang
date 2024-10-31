package main

import (
	"fmt"
)

func main() {

	var fruits [4]string

	fruits = [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println(fruits)

	fruits = [4]string{
		"Apple",
		"Banana",
		"Orange",
	}

	fmt.Println(fruits)


	var numbers= [...]int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(numbers)
	fmt.Println(len(numbers))

 	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
			fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits {
			fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	fruits = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruits {
			fmt.Printf("nama buah : %s\n", fruit)
	}

	// kalo mau buat array dengan make() variabelnya harus slice dan belum di inisialisasi
	var test = make([]string, 4)
	test[0] = "apple"
	test[1] = "manggo"

	fmt.Println(test)
}
