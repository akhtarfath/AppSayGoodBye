package main

import (
	"fmt"
	sayGoodBye "github.com/akhtarfath/SayGoodBye"
	sayGoodBye2 "github.com/akhtarfath/SayGoodBye/v2"
)

func main() {
	fmt.Println(sayGoodBye.SayGoodBye())
	fmt.Println(sayGoodBye2.SayGoodBye("Muhammad A. Fathan"))
}
