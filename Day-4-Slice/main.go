package main

import "fmt"

func main() {
	fmt.Println("main function is started here")

	mySlice()
	fmt.Println("mySlice function is finished here")
	mySliceAdv()

}

func mySlice() {
	fmt.Println("mySlice is started here")

	Universal := []string{"Sun", "Moon", "Earth", "Planet"}
	fmt.Println(len(Universal))
	fmt.Printf("The value of Universal %q\n", Universal)
	Universal[2] = "Hell"
	fmt.Println(Universal[2])

	for i := 0; i < len(Universal); i++ {
		fmt.Println(Universal[i])

	}
}

func mySliceAdv() {
	fmt.Println("mySliceAdv is started here")
	arr := [5]int{100, 1000, 10000, 100000, 1000000}
	fmt.Println("Print array values: ", arr)

	slice1 := arr[2:4]
	fmt.Println("Print slice values: ", slice1)
	slice1 = append(slice1, 100, 1000)
	fmt.Println("Print slice values: ", slice1)
}
