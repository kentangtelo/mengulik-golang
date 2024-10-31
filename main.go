package main

import "fmt"


func main()  {
	var firstName string = "John";

	var lastName string;
	lastName = "Doe";

	thirdName := "Smith";

	fmt.Printf("halo %s %s %s!\n", firstName, lastName, thirdName)
	fmt.Println("halo", firstName, lastName, thirdName + "! \n")

	thirdName = "Wick";

	fmt.Printf("halo %s %s %s!\n", firstName, lastName, thirdName)
	fmt.Println("halo", firstName, lastName, thirdName + "! \n")

}
