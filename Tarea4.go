package main

import (
	"fmt"
	"strings"
)

func main() {
    
	fmt.Println("Ejercicio 3: Hallar el mayor valor de una lista")
    Lista := []int{}
	var numero int32

	for {
		fmt.Println("Ingrese numeros enteros para añadir a la lista:")
		fmt.Scanln(&numero)

		Lista = append(Lista, int(numero))

		var respuesta string
		fmt.Println("Desea añadir mas numeros (S/N)")
		fmt.Scanln(&respuesta)

		if strings.ToUpper(respuesta) == "N" {
			break
		}else{
			continue
		}
	}
	
	fmt.Println("la lista:",Lista)

	mayor := Lista[0]
	
    for _, numerocomparacion := range Lista {
        if numerocomparacion > mayor {
            mayor = numerocomparacion
        }
    }

	fmt.Println("El numero mayor de la lista:", mayor)
  
}

