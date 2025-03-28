package main

import "fmt"

func main() {
	a, b := 10, 3
	fmt.Println(" A Soma é: ", a + b)    
	fmt.Println(" A Subtração é: ", a - b)
	fmt.Println(" A Divisão é: ", a / b)
	fmt.Println(" A Multiplicação é: ", a * b)
	fmt.Println(" O Resto da Divisão é: ", a % b)

	a+=1
	fmt.Println("Incrementar a", a)

	if a > 0 && b > 0 {
		fmt.Println("Numeros Positivos")
}	

	}
