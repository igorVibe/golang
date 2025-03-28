package main

import "fmt"

func main() {
	var a int;
	var b int;

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&a)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&b)

	fmt.Println("A Soma Dos Numeros é: ", a + b)
	fmt.Println("A Subtração é: ", a - b)
	fmt.Println("A Divisão é: ", a / b)
	fmt.Println("A Multiplicação é: ", a * b)
}
