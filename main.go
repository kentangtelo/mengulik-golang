package main

import (
	"fmt"
)

func main() {
	const (
		square = "kotak"
		isTodayFriday = true
		numeric uint8 = 255
		floatNum = 2.2
	)

	const (
		a = "test"
		b
	)

	const (
    today string = "senin"
    sekarang
    isToday2 = true
	)

	fmt.Println(square)
	fmt.Println(isTodayFriday)
	fmt.Println(numeric)
	fmt.Println(a)
	fmt.Println(b)

	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"

	fmt.Println(satu)
	fmt.Println(dua)
}
