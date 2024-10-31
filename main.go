package main

import (
	"fmt"
)

func main() {
	var xs ="abcdefghijklmnopqrstuvwxyz"

	for i, v := range xs {
		fmt.Println(i, string(v))
	}

	var ys = [5]int{1, 2, 3, 4, 5}

	for _, v := range ys {
		fmt.Println(v)
	}

	var zs = ys[1:3]

	for _, v := range zs {
		fmt.Println(v)
	}

	var kvs = map[byte]int{'a': 1, 'b': 2, 'c': 3}
	for k, v := range kvs {
		fmt.Println(string(k), v)
	}

	for range kvs {
		fmt.Println("Done")
	}

	for i:= range 5{
		fmt.Println(i)
	}

}
