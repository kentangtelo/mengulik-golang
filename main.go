package main

import (
	"fmt"
)

func main() {

	var value = ((2+2) * 2) / 2
	var isequal = value == 4

	fmt.Printf("value: %d %t\n", value,isequal)

	var point1 = 4

	if point1 > 7 {
		fmt.Println("Good")
	} else if point1 > 5 {
		fmt.Println("Not Bad")
	} else {
		fmt.Println("So Bad")
	}

	var point2 = 6555.0

	if percent := point2 / 100; percent >= 100 {
		fmt.Printf("%.1f%s Perfect!\n", percent, "%")

	}else if percent >= 70  {
		fmt.Printf("%.1f%s Good\n", percent, "%")

	}else {
		fmt.Printf("%.1f%s Not Bad\n", percent, "%")
	}

	var point = 64

	switch point {
		case 100:
			fmt.Println("Perfect!")
		case 70:
			fmt.Println("Good")
		case 50,64:
			fmt.Println("Not Bad")
		default:
			fmt.Println("Not Bad")
	}

	point = 6

	switch {
	case point == 8:
			fmt.Println("perfect")
	case (point < 8) && (point > 3):
			fmt.Println("awesome")
	default:
			{
					fmt.Println("not bad")
					fmt.Println("you need to learn more")
			}
	}

	point = 6

	switch {
	case point == 8:
			fmt.Println("perfect")
	case (point < 8) && (point > 3):
			fmt.Println("awesome")
			fallthrough
	case point < 5:
			fmt.Println("you need to learn more")
	default:
			{
					fmt.Println("not bad")
					fmt.Println("you need to learn more")
			}
	}


 point = 5

	if point > 7 {
			switch point {
			case 10:
					fmt.Println("perfect!")
			default:
					fmt.Println("nice!")
			}
	} else {
			if point == 5 {
					fmt.Println("not bad")
			} else if point == 3 {
					fmt.Println("keep trying")
			} else {
					fmt.Println("you can do it")
					if point == 0 {
							fmt.Println("try harder!")
					}
			}
	}

}
