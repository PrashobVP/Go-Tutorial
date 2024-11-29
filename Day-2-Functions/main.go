package main

import "fmt"

func main1() {
	fmt.Println("VisualPath Day 2")

	//conditional statement
	//example 1
	// if 10> 5 is true then it will execute inside block statement .otherwise it will go to second condition and execute block.
	if 10 > 5 {
		fmt.Println("10 is greater than 5")
	} else {
		fmt.Println("10 is not greater than 5")
	}
    var returnValue = myFunc("my 1st var aruguments")
	fmt.Println(returnValue)

	fmt.Println("2nd function ended here")

	myloop()
	fmt.Println("for loop ended here")

	MyIfElse()
	
}

func myFunc(myVar string) int{
	fmt.Println("2nd function started here")
	fmt.Println("myVar value is :", myVar)

	return 100
}

//for loop

func myloop() {
	fmt.Println("for loop started here")

	for i := 0; i < 5; i++ {
		fmt.Println("My name is Prashob in loop", i)
	}
}

func MyIfElse() {
	fmt.Println("if else started here")

	fmt.Println("Do you have a mobile")
	haveMobile := true
	phNo:= 9500073161
	if haveMobile == true {
		fmt.Println("I have a mobile",phNo)
		
	} else {
		fmt.Println("I don't have a mobile")
	}
}
