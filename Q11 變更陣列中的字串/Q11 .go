package main

import (
	"fmt"
)

func main() {
	var cars = [4]string{"Toyota", "Yamaha", "Benz", "BMW"}
	cars[0] = "off"
	fmt.Print(cars)
}
