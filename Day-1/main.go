package main

import "fmt"

func main() {

	fmt.Println("VisualPath Day 1")
	//Datatype String

	var name string = "Prashob"
	fmt.Println("My name is ", name)
    
	// datatype integer
	var age int = 32
	DOB := "28-01-1990"
    var mobileNumber int = 1234567890
	fmt.Println("My age is ", age , "and DOB is ",DOB)
	fmt.Println("My mobile number is ", mobileNumber)

	// datatype float
	var height float32 = 5.5
	var width float64 = 100.5

	fmt.Println("My height is ", height, "and width is ", width)

	//Datatype Boolean
	var isMarried bool  //go will take default value as false
	fmt.Println("Married Status ", isMarried)

	//Performing arithmetic operations
	var a int = 10
	var b int = 20

	c := a + b

	fmt.Println("Sum of a and b is ", c)

	var x int = 100
	var y int = 2000

	var d int = x * y
	var e int = x / y
	var f int = y % x- y

	fmt.Println("Multiplication of x and y is ", d)
	fmt.Println("Division of x and y is ", e)
	fmt.Println("Modulus of x and y is ", f)

      //modify the program and use varuables to get the same output
	  var g int = 10
	  var h int = 20
	  fmt.Println("Sum of g and h is ", g+h)
      
	  //modify the variables to get a new output

	  g = 100
	  
	  fmt.Println("Modifed the value of g is ", g)
}
