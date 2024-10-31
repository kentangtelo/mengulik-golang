package main

import "fmt"


func main()  {
	var firstName string = "John";

	var lastName string;
	lastName = "Doe";

	thirdName := "Smith";

	var first, second, third string;
	first, second, third = "John", "Doe", "Smith";

	var fourth, fifth, sixth string = "John1", "Doe1", "Smith1";

	seventh, eighth, ninth := "John2", "Doe2", "Smith2";

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello";

	name := new(string)
	boolEan := new(bool)
	intEger := new(int)

	fmt.Println("name string", name);
	fmt.Println("boolEan bool", boolEan);
	fmt.Println("intEger int", intEger);

	fmt.Println(firstName, lastName, thirdName);
	fmt.Println(first, second, third);
	fmt.Println(fourth, fifth, sixth);
	fmt.Println(seventh, eighth, ninth);
	fmt.Println(one, isFriday, twoPointTwo, say);

	fmt.Printf("halo %s %s %s!\n", firstName, lastName, thirdName)
	fmt.Println("halo", firstName, lastName, thirdName + "! \n")

	thirdName = "Wick";

	fmt.Printf("halo %s %s %s!\n", firstName, lastName, thirdName)
	fmt.Println("halo", firstName, lastName, thirdName + "! \n")

}
