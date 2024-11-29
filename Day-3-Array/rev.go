package main

import "fmt"

func main() {

	fmt.Println("main function start here")
	myArr()
	fmt.Println("myArr function end here")
	fmt.Println("main function end here")
}

func myArr() {
	fmt.Println("myArr start here")
	coral := [5]string{"blue coral", "staghorn coral", "pillar coral", "elkhorn coral"}

	fmt.Printf("The value of coral %q\n", coral) //format print with %q means put quotation
	fmt.Println(coral[2])                        //index 2 value
	//how to change index 2 value to Prashob coral

	coral[2] = "Prashob coral"

	fmt.Printf("The value of coral %q\n", coral)
	fmt.Println(coral[2])
	fmt.Println("the count of array", len(coral))
	//coral = append(coral, "black coral")
	coral[4] = "hello world"
	fmt.Printf("The value of coral %q\n", coral)
}
