package main

import "fmt"

func main() {

	fmt.Println("Welcome to Day 3 with Array")

	myArrayTest()
}

func myArrayTest() {

	arr1 := [9]int{1, 2, 3, 4, 5, 10, 20, 30, 40}
	fmt.Println("my array value for arr1", arr1)
	arr1[0] = 1
	fmt.Println("my array value updated with 0 index for arr1", arr1[0])

	arr1[4] = 100
	fmt.Println("my array modified value for arr1", arr1)

	arr1[3] = arr1[1] + arr1[2]

	fmt.Println("my array modified value for arr1", arr1)

	fmt.Println("length of my arr1", len(arr1))
	total := 0
	for i := 0; i < len(arr1); i++ {

		fmt.Printf("arr1[%d] = %d\n", i, arr1[i])
		total = total + arr1[i] //0+[0] index from array
		// even and odd
		if i%2 == 0 {
			fmt.Println(arr1[i], "even")
		} else {
			fmt.Println(arr1[i], "odd")
		}
	}
	fmt.Println("total", total)

}
