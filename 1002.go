package main

import (
	"fmt"
)

func main() {
	var raio float64
	var area float64
	fmt.Scanf("%f", &raio)
	area = ((raio * 2) * 3.14159)
	fmt.Println("A= ", area)
}
