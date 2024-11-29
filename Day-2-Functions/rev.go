package main

import "fmt"

func main() {

	fmt.Println("main block start here")

	myIfElse()
	fmt.Println("my if block is end here")

	fmt.Println("main again started here")
	myLoop()
	fmt.Println("main loop end here")
	fmt.Println("main again started here")
	c, d := myFunction("hai", 2024)
	fmt.Println(c, d)

	fmt.Println("myFUNC is starting here")
	e, f := myfunc("Prashob")
	myfunc("a", "b")

	fmt.Println(e, f)
	fmt.Println("main function end here")
}

func myFunction(myVar string, myint int) (string, string) {
	fmt.Println("Input string value is ", myVar, "and", "Input int value is ", myint)
	fmt.Println("myFunction begin here")
	return "hello world", "Golang"
}

func myfunc(name ...string) (arg1 int, arg2 int) {

	fmt.Println("name is ", name)
	arg1 = 100
	arg2 = 200
	return arg1, arg2
}

func myIfElse() {

	fmt.Println("my if block will start here")

	a := 5

	if a > 5 {

		fmt.Println("a is greater than 5")
	} else if a < 5 {
		fmt.Println("a is less than 5")
	} else {
		fmt.Println("a is equal to 5")
	}
}

func myLoop() {
	b := "Prashob is master in GO"
	for i := 0; i < 10; i++ {

		fmt.Println(b)
	}
}
