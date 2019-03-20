package main

import "fmt"

func main() {
	numeros := [...] int {18, 28, 21, 96, 97}
	numeroMayor := numeros[0] //Asignar variable al primer elemento
	for _, numero := range numeros {
		if numero > numeroMayor{
			numeroMayor = numero
		}
	}
	fmt.Printf("El número mayor entre %v es %d", numeros, numeroMayor)
}