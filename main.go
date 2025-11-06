package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{20,25,30}
	// var newAges=[3]int{20,25,30}
	// fmt.Print(ages,newAges)

	// // fmt.Println(ages, len(ages))
	names:= [4]string{"Yoshi", "Som", "Amrut", "Vlad"}
	// fmt.Println(names, len(names))
	
	//Slices(use arrays under the hood)
	// var scores = []int{100,50, 60}

	// fmt.Println(scores);
	// scores=append(scores,85)
	// fmt.Println(scores, len(scores));
	
	//slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)
	
	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

}
