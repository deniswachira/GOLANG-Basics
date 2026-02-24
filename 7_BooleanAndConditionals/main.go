package main

import "fmt"

func main(){
	//boolean and conditionals
	fmt.Println("=====Boolean and Conditionals=====")

	age := 35
	// if conditional
	if age > 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// if else if
	if age < 18 {
		fmt.Println("You are a minor")
	} else if age >= 18 && age < 60 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a senior citizen")
	}
}