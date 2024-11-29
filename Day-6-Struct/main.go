package main

import "log"

type Circle struct {
	Radius int
}
type Rectangle struct {
	Length int
	Bredth int
}

func main() {
	log.Println("main function")
}

func () myStruct() {
	log.Println("Struct starting")


	var mySmallCircle Circle
	mySmallCircle.Radius = 3
	mySmallRectangle := Rectangle{
		Length: 3,
		Bredth: 2,
	}
	log.Println("radius of the circle is ", mySmallCircle.Radius)
}

func () Name(){

}
