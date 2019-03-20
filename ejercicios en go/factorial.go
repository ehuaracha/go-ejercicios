package main

import "fmt"

func main() {
   // factorial(5) = 5 * 4 * 3 * 2 * 1 = 120
   fmt.Println("Factorial de 5: ", factorial(5))
}

// Funcion que calcula un factorial
func factorial(n int) int {
   if n == 0 {
      return 1
   }

   return n * factorial(n-1)
}