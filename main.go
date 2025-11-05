package main

import "fmt"

func main() {

	age := 23
	name := "Amrutanshu"
	//Print
	fmt.Print("hello, ")
	fmt.Print("world!\n")
	fmt.Print("new line\n")

	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas!")
	// fmt.Println("my age is",age)

	//formatted string same as C language use %v to directly get the variable data types
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T\n", age);
	fmt.Printf("you scored %0.1f points!\n", 10.12);
	
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Printf(str)
}
