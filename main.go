package main

import (
	"fmt"
)


func dadosPessoa(nome string, idade int) (int, string) {
	if idade >= 18 {
		return idade, "maior de idade"
	} else {
		return idade, "menor de idade"
	}
}

func main() {
	nome := "Igor"
	idade := 16

	id, status := dadosPessoa(nome, idade)

	fmt.Printf("%s tem %d anos e é %s.\n", nome, id, status)
}



	


	
	

