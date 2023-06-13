package main

import (
    "fmt"
)

func plusMinus(arr []int32) {
    inputArr := arr
    var positiveNumbers float32
    var negativeNumbers float32
    var zeroNumbers float32
    for i,_ := range inputArr{
        if inputArr[i] > 0 {
            positiveNumbers = positiveNumbers + 1
        } else if inputArr[i] < 0 {
                negativeNumbers = negativeNumbers + 1
        } else {
            zeroNumbers = zeroNumbers + 1
        }
            
    }
    length := float32(len(arr))
    
    fmt.Printf("%.6f\n", positiveNumbers/length)
    fmt.Printf("%.6f\n", negativeNumbers/length)
    fmt.Printf("%.6f\n", zeroNumbers/length)
}

func main() {
    arr := []int32{-1, 2, 0, 5, -3, 0, 1}
    plusMinus(arr)
}
