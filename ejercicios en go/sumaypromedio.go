package main
 
import "fmt"
 
func main() {
	var calificacionIntroducida int // Para escanear cada calificacion introducida por el usuario
	var longitudArreglo int //La longitud
 
	/*
		Preguntar la longitud...
	*/
	fmt.Print("¿Cuántas calificaciones desea promediar? ")
	fmt.Scanf("%d\n", &longitudArreglo)
 
	/*
		En calificaciones guardaremos los valores
		Como no podemos asignar una longitud a un array en tiempo de ejecución, usamos
		un slice. Para crearlo llamamos a make y pasamos como primer argumento
		el tipo de dato que sería int, y como segundo la longitud que
		fue previamente escaneada
	*/
	calificaciones := make([] int, longitudArreglo) //Aquí guardaremos los valores.
 
	/*
		Preguntar cada valor...
	*/
	for i:= 0; i < longitudArreglo; i++{
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