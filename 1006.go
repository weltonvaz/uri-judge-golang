package main

import "fmt"

func main() {

	var A, B, C float64
	fmt.Scanln(&A)
	fmt.Scanln(&B)
	fmt.Scanln(&C)
	s := fmt.Sprintf("%.1f", ((2*A)+(3*B)+(5*C))/10)
	fmt.Println("MEDIA =", s)
}
