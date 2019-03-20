package main
import "fmt"
func fibonacci(i int) int {
  //Caso base
  if(i == 0) {
    return 0
  }
  //Otro caso base
  if(i == 1) {
    return 1
  }
  return fibonacci(i-1) + fibonacci(i-2)
}
func main() {
  for n := 0; n < 15; n++ {
    fmt.Printf("| %d ", fibonacci(n))
  }
  fmt.Printf("|\n")
}