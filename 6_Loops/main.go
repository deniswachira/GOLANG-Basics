package main

import "fmt"

func main(){
	//loops
	fmt.Println("=====Loops=====")

	// x := 0
 	// for init; condition; post{}
	// for x < 10 {
	// 	fmt.Println("Value of x is", x)
	// 	x++
	// }

	names := []string{"Denis", "Mark", "Owen", "Dan"}
	// for loop
	for i := 0; i < len(names); i++ {//i is the index of the element in the array
		fmt.Printf("index is %v and name is %v \n", i, names[i])
	}
	//range loop
	fmt.Println("====Range Loop====")
	for index, name := range names { //index is the index of the element and name is the value of the element
		fmt.Printf("index is %v and name is %v \n", index, name)
	}
}