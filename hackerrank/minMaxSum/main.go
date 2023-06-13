package main

import (
    "fmt"
    "sort"
)

func miniMaxSum(arr []int32) {
    // Write your code here
    inputArr := arr
   
    sort.Slice(inputArr, func(i,j int) bool {
        return inputArr[j] > inputArr[i]
    })
    var minSum int
    var maxSum int
    
    for i := 0; i < (len(inputArr) - 1); i++ {
        minSum += int(inputArr[i])
    }
    
    for i := 1; i < len(inputArr); i++ {
        maxSum += int(inputArr[i])
    }
    
   fmt.Println(minSum, maxSum)
}

func main() {
    arr := []int32{9, 3, 6, 1, 8}
    miniMaxSum(arr)
}

