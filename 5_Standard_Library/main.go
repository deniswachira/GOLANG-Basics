package main

import (
	"fmt"
	"sort"
)

func main() {
	//standard library
	fmt.Println("=====Standard Library====")

	// text := "This is string with a lot of words and characters"

	// fmt.Println(strings.Contains(text, "string"))
	// fmt.Println(strings.Count(text, "a"))
	// fmt.Println(strings.HasPrefix(text, "This"))
	// fmt.Println(strings.HasSuffix(text, "characters"))
	// fmt.Println(strings.Index(text, "a"))
	// fmt.Println(strings.LastIndex(text, "a"))
	// fmt.Println(strings.ReplaceAll(text, "a", "A"))
	// fmt.Println(strings.ToUpper(text))
	// fmt.Println(strings.ToLower(text))

	ages := []int{80, 30, 40, 70, 50, 60, 70}

	fmt.Println("Ages before sorting", ages)
	//sort
	// sort.Ints(ages)
	// fmt.Println("Ages after sorting", ages)

	//searching
	index := sort.SearchInts(ages, 50)
	fmt.Println("Index of 50 is ", index)

	names := []string{"Denis", "Mark", "Owen", "Dan"}
	fmt.Println("Names before sorting", names)
	sort.Strings(names)
	fmt.Println("Names after sorting", names)
}
