package main

import "fmt"

func main() {
	numeros := [...] int {18, 28, 21, 96, 97, 30, 50, 128, -1, -50, 55, 800, 450}
	numeroMayor, numeroMenor := numeros[0], numeros[0]
	for _, numero := range numeros {

		if numero > numeroMayor{
			numeroMayor = numero
		}

		if numero < numeroMenor{
			numeroMenor = numero
		}
	}
	fmt.Printf("En el arreglo %v\nEl número menor es %d. \nY el mayor es %d\n", numeros, numeroMenor, numeroMayor)
}