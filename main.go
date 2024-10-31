package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i%2==1 {
			continue
		}

		if i>8 {
			break
		}

		fmt.Println("angka", i)
	}


	fmt.Println()
	/*
	*****
	****
	***
	**
	*
	*/

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	fmt.Println()
	/*
	*
	**
	***
	****
	*****
	*/

	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	/*
	*
	**
	***
	****
	*****
	****
	***
	**
	*
	*/

	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			if i < 5 {
				if j <= i {
					fmt.Print("*")
				}
			} else {
				if j < 8-i {
					fmt.Print("*")
				}
			}
		}
		fmt.Println()
	}

	fmt.Println()

	/*
	*****
	 ****
	  ***
		 **
		  *
	*/

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j >= i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Println()

	/*
	    *
	   **
	  ***
	 ****
	*****
	*/

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j >= 4-i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Println()

	/*
	    *
	   **
	  ***
	 ****
	*****
	 ****
	  ***
		 **
		  *
	*/

	for i := 0; i < 9; i++ {
		for j := 0; j < 5; j++ {
			if i < 5 {
				if j >= 4-i {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			} else {
				if j >= i-4 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}

	/*
	    *
	    **
	   ****
	  ******
	 ********
	**********
	 ********
	  ******
		 ****
		  **
			*
	*/

	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if i < 6 {
				if j >= 5-i && j <= 5+i {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			} else {
				if j >= i-5 && j <= 15-i {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}

	fmt.Println()

}
