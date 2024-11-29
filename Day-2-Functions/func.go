package main

import "fmt"

func main() {
	fmt.Println("My main function start here")

	fmt.Println("The returnValue will be start here")
	//returnValue("hello")
	fmt.Println(returnValue("hello"))

	//2nd fucntion start here
	fmt.Println("The returnMultiValue will be start here")
	fmt.Println(returnMultiValue(300))


}

func returnValue(myvar string) string {
	fmt.Println("The value of myvar", myvar)
	return "world"

}

func returnMultiValue(Totalmark int) (int, int) {
	fmt.Println(Totalmark)
	return 10,20
}