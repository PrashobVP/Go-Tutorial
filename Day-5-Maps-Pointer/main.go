package main

import "fmt"

func main() {
	fmt.Println("Pointer and Maps start here")
	myMap()
	myPointer1()
}

func myPointer1() {
	fmt.Println("Pointer start here")
	var p *int
	x := 1000
	p = &x
	fmt.Println("Value stored in x = ", x)
	fmt.Println("address of x = ", &x)
	fmt.Println("Value of p", p)

}

func myMap() {
	a := map[string]int{"prashob": 39, "minnu": 34}

	fmt.Println("Print value of Map", a["prashob"])

	b := map[string]int{}
	fmt.Println("Print value of Map", b)
	a["prashob"] = 40
	for _, v := range a {

		fmt.Println(v)
	}
}

/*func myPointer() {

} */
