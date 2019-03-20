package main
 
import "fmt"
 
func main() {
	var calificacionIntroducida int // Para escanear cada calificacion introducida por el usuario
	var calificaciones[10] int //Aquí guardaremos los valores
 
	/*
		Preguntar cada valor...
	*/
	for i:= 0; i < 10; i++{
		fmt.Printf("\nIntroduce la calificación %d: ", i + 1)
		fmt.Scanf("%d\n", &calificacionIntroducida)
		calificaciones[i] = calificacionIntroducida // y asignarlo
	}
 
	/*
		Calcular promedio
	*/
 
	sumatoria := 0
	for _,calificacion := range calificaciones{
		sumatoria += calificacion
	}
 
	promedio := sumatoria / len(calificaciones)
	fmt.Printf("El promedio es %d", promedio)
}