package main

import "fmt"

func main() {
	var usuario string
	var senha string
   
	fmt.Print("DIgite o Usuario ")
	fmt.Scanln(&usuario)

	fmt.Print("Digite a Senha")
	fmt.Scanln(&senha)


	if usuario == "admin" && senha == "2008"{
	fmt.Println("Login Feito com Sucesso")

	}else{ 
		fmt.Println("Login Errado")
	
   

	
 }
	
	
}
