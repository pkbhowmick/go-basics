package main

import "fmt"


func integerToBinary(number int) string {
	var binary string = ""
	for number!=0 {
		if number%2 == 1 {
			binary = "1" + binary
		} else {
			binary = "0" + binary
		}
		number/=2
	}
	return binary
}

func main(){
	var number int
	fmt.Print("Enter integer number: ")
	fmt.Scanf("%d",&number)
	fmt.Println(integerToBinary(number))
}