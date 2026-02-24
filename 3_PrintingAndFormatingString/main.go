package main

import "fmt"

func main (){
	//print
	fmt.Println("====Formatting strings 😂 ====")
	// fmt.Print("World: \n")
	// fmt.Print("new line \n")

	// fmt.Println("Hello Denis")
	// fmt.Println("GoodBye Denis")


	myAge := 35
	myFirstName :="Denis"

	fmt.Println("My age is ", myAge, "and my name is", myFirstName)

	//formating string %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n",myAge,myFirstName)

	fmt.Printf("my age is %q and my name is %q \n", myAge,myFirstName)

	//getting the type
	fmt.Printf("my age is of type %T \n", myAge)

	// out putting floats
	fmt.Printf("you scored %f points! \n", 255.46)
	fmt.Printf("you scored %0.2f points! \n", 255.46)


	//sprintf(save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", myAge, myFirstName)
	fmt.Println("Saved string is:", str)
}