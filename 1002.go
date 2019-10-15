package main

import "fmt"

func main() {

	var raio float64
	var resul float64
	pi := 3.14159
	fmt.Scanln(&raio)
	fmt.Printf("A=")
	resul := raio * *2
	fmt.Printf("%.4g", pi*resul)
}
