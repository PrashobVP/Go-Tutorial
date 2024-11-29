package main

import "fmt"

func main() {
	fmt.Println("====Home assignment- Day2===")

	tSum, tDif := Sum(200, 4)
	fmt.Println("Total of SUM = ", tSum, "and", "Total of Dif = ", tDif)

	numPrint()

	fmt.Println("back to main function")
}

//write a function that takes two integers and reruns the sum and the difference

func Sum(a, b int) (int, int) {
	fmt.Println("The Input value of a is", a, "and", b)
	fmt.Println("====Exercise 1=======")
	fmt.Println("Starting arithmetic function,(a+b) && (a-b) !!!")
	sum := a + b
	dif := a - b

	return sum, dif
}

//Write a program that prints number from, 1 to 10 and prints "even" or "odd" next to each

func numPrint() {
	fmt.Println("=====Exercise 2======")
	fmt.Println("Starting even & Odd function")

	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			fmt.Println(i, "Even")
		} else {
			fmt.Println(i, "Odd")

		}
	}
}
