package main

import "fmt"

func main() {

	var A, B float64
	fmt.Scanln(&A)
	fmt.Scanln(&B)
	fmt.Println("MEDIA =", ((3.5*A)+(7.5*B))/(3.5+7.5))
}
