package main

import "fmt"

func main(){
	//arrays
	fmt.Println("===== Arrays =====")

	// var marks [3]int = [3]int{30,40,49}

	//shorthand
	var marks = [3]int{30,35,89}

	fmt.Println("Marks", marks)
	// length
	fmt.Println("Length of the array is ", len(marks))

	names := [4]string{"denis","mark","Owen","Dan"}

	fmt.Println(names)

	// SLices [uses array under the hook]
	fmt.Println("====Slices======")

	var scores = []int{40,90,30}
	scores[2] = 40
	scores = append(scores,85)

	fmt.Println(scores,len(scores))

	//slice range
	rangeOne := names[1:3]
	rangeTwo := names[2:] //from index 2 to the end of the array
	rangeThree := names[:3] //from the start of the array to index 3

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)


}