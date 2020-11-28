package main

import (
	"fmt"
	"strconv"
)

func main(){
	var n int = 100

	// Converting integer to string
	var str string = strconv.Itoa(n)
	fmt.Printf("%s\n",str)

	//Converting string to integer
	convertedInt, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error in convertion from string to int")
	} else {
		fmt.Printf("%d\n",convertedInt)
	}

	convertedInt, err = strconv.Atoi("abcd")            // will give error because "abcd" can't be converted to int
	if err != nil {
		fmt.Println("Error in convertion from string to int")
	} else {
		fmt.Printf("%d\n",convertedInt)
	}


}