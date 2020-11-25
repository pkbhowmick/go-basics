package main

import "fmt"

var fib [105]int

func printNthFibonacci(n int) int {
	if n==1 {
		return 0
	}
	if n==2 {
		return 1
	}
	if fib[n]!=0 {
		return fib[n]
	}
	fib[n] = printNthFibonacci(n-1)+printNthFibonacci(n-2)
	return fib[n]
}


func main(){
	var n int
	fmt.Scanf("%d",&n)
	if n>90 {
		fmt.Println("Fibonacci of the given value is too big!")
	} else {
		fmt.Println(printNthFibonacci(n))
	}
}