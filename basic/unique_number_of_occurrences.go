package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
    mp := make(map[int]int)
    
    for i:=0;i<len(arr);i++ {
        mp[arr[i]]++
    }
    
    mp2 := make(map[int]int)
    
    for _,val := range mp {
        mp2[val]++
    }
    
    return (len(mp)==len(mp2))
}

func main() {
	arr := []int{1,2,2,1,1,3}
	fmt.Println(uniqueOccurrences(arr))
}