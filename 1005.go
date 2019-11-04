package main

import "fmt"

func main() {

	var A, B float64
	fmt.Scanln(&A)
	fmt.Scanln(&B)
	s := fmt.Sprintf("%.5f", ((3.5*A)+(7.5*B))/(3.5+7.5))
	fmt.Println("MEDIA =", s)
}
