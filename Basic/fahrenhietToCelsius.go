package main

import "fmt"

func fahrenhietToCelsius(f float64) float64{
	var c float64 = (f-32.0)*5.0/9.0
	return c
}

func main(){
	var fahrenhiet float64 
	fmt.Print("Enter fahrenhiet: ")
	fmt.Scanf("%f",&fahrenhiet)
	fmt.Println(fahrenhietToCelsius(fahrenhiet))
}