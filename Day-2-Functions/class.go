package main

import "fmt"

func main() {

	fmt.Println("Welcome to India")
	myFunc("hello world")
	myMark(68)
	a, b := myReturn("Prashob Preman", 39)

	fmt.Println("I am back in the main code with a value", a, b)
}

//this is a sample function with no argument and no returns

func myFunc(myVar string) {
	fmt.Println("Input argument to my function is ", myVar)
	fmt.Println("myFunc starting here")
	//loop

	for i := 0; i < 5; i++ {
		fmt.Println(myVar, "the value of i", i)

	}

}

func myMark(myTotalMark int) {

	fmt.Println("Input argument to my function is ", myTotalMark)
	fmt.Println("myMark starting here")

	if myTotalMark > 75 {
		fmt.Println("You have a distiction")
	} else if myTotalMark > 60 {
		fmt.Println("You have a 1st class")
	} else if myTotalMark > 50 {
		fmt.Println("You have a 2nd class")
	} else {
		fmt.Println("You failed ")
	}
}

func myReturn(Name string, Age int) (arg1 int, arg2 int) {
	arg1 = 101
	arg2 = 20
	fmt.Println("Input argument to my function is ", Name, "and my Age:", Age)
	fmt.Println("myReturn starting here")
	return arg1, arg2

}
