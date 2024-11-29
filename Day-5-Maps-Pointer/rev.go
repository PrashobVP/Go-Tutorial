package main

import "fmt"

func main() {
	fmt.Println("Main code is starting here")
	phMapF()
	fmt.Println("Main code is starting here")
	myPointer("hello")

}

func phMapF() {
	hasPhoneNumber := map[string]bool{"Prashob": true, "Shivani": true, "Minnu": true, "Sriyaan": false}
	phMap := map[string]int{"Prashob": 9500073161, "Shivani": 5342344489, "Minnu": 5342354773, "kanchi": 95672523558}
	finalMap := make(map[string]int)
	for k, v := range hasPhoneNumber {
		name := k
		hasPh := v

		if hasPh {
			phValue := phMap[name]
			fmt.Printf("%s has ph num as : %d \n", name, phValue)
			finalMap[name] = phValue
		}

	}
	fmt.Println("The output map for home assigment is : ", finalMap)

}

func myPointer(input string) {

	var p *int

	i := 400

	p = &i
	fmt.Println("the value of i", i)
	fmt.Println("the value of *p", *p)

	i = 100
	fmt.Println("the value of i", i)
	fmt.Println("the value of *p", *p)

	*p = *p + 100
	fmt.Println("the value of i", i)
	fmt.Println("the value of *p", *p)
	fmt.Println("the input value ", input)

}
