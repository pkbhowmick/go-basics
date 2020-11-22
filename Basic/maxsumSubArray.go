package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
    maxVal := -math.MaxInt32
    sum := 0
    for i:=0; i<len(nums);i++ {
        sum += nums[i]
        if sum > maxVal {
            maxVal = sum
        }
        if sum < 0 {
            sum = 0
        }
    }
    return maxVal
}

func main() {
	arr := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(arr))
}