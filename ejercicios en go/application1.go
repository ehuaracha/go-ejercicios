package main

import "fmt"

func main() 
    int []arreglo=new int [10];
    for (int i=1; i<6; i++){
    System.out.print("Escriba el elemento: "+i+" ");
    arreglo[i]=sc.nextInt();
    }
    int mayor=0, menor=0, promedio=0;
        for (int i=1; i<6; i++){
        promedio=arreglo[i]+promedio;
        if (arreglo[i]>mayor){
        mayor=arreglo[i];
        }
        if(arreglo[i]<menor){
        menor=arreglo[i];}
        }
        System.out.println("Nro mayor: "+mayor);
        System.out.println("Nro menor: "+menor);
        System.out.println("promedio: "+promedio/5);
    }       
}