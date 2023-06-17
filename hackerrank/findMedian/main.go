package main

import (
    "fmt"
    "sort"
)



/*
 * Complete the 'findMedian' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func findMedian(arr []int32) int32 {
    // Write your code here
    sort.Slice(arr, func(i,j int) bool {
        return arr[j] > arr[i]
    })
    middleIndex := len(arr) / 2
    middleElement := arr[middleIndex]
    return middleElement
    

}

func main() {
    arr := []int32{4,5,8,3,5,1,10,11,2}
    result := findMedian(arr)
    fmt.Println(result)

}

