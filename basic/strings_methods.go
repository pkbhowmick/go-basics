package main

import (
	"fmt"
	"strings"
)

func main(){
	var str string = "This is an example of string"

	// HasPrefix example: strings.HasPrefix(str, prefix) bool
    var res bool = strings.HasPrefix(str,"This")
	fmt.Printf("%t\n",res)

	//HasSuffix example: strings.HasSuffix(str, suffix) bool
	res = strings.HasSuffix(str, "ring")
	fmt.Printf("%t\n",res)

	//Contains example: strings.Contains(string, substring) bool
	res = strings.Contains(str, "example")
	fmt.Printf("%t\n",res)

	//Index returns first match index of substring. example: strings.Index(string, substring) int 
	var pos int = strings.Index(str,"is")
	fmt.Printf("%d\n",pos)

	//LastIndex returns last match index of sunstring. example: strings.LastIndex(string, sunstring) int
	pos = strings.LastIndex(str,"is")
	fmt.Printf("%d\n",pos)

	//Replaceing substring. example: strings.Replace(str,old substr,new substr,n) string, n for first n occurences, n=-1 for all occarences
	var newString string = strings.Replace(str,"example","new example",-1)
	fmt.Printf("%s\n%s\n",str,newString)

	//Count, counting occurences of a substring. example: strings.Count(str, substr) int
	var cnt int = strings.Count(str,"is")
	fmt.Printf("%d\n",cnt)

	//Changing the case of the string
	fmt.Printf("%s\n",strings.ToLower(str))
	fmt.Printf("%s\n",strings.ToUpper(str))

	//TrimSpace used for trimming trailing or leasding space and "Trim" is used for trimming specific string
	fmt.Printf("%s\n",strings.Trim(str,"string"))

	//strings.Fields used for splitting string on only whitespaces and strings.Split for separator
	var arr []string = strings.Fields(str)
	for _,item := range arr {
		fmt.Printf("%s\n",item)
	}
	arr = strings.Split(str," ")
	for _,item := range arr {
		fmt.Printf("%s\n",item)
	}

}