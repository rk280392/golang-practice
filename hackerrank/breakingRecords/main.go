package main

import (
    "fmt"
)

func breakingRecords(scores []int32) []int32 {
    // Write your code here
    
    
    inputArr := scores
    if len(inputArr) == 0 {
        return []int32{}
    }
    var minMax []int32
    minScore := inputArr[0]
    maxScore := inputArr[0]
    var countMax int32 
    var countMin int32
    for i := 1; i < len(inputArr); i++ {
        if inputArr[i] < minScore {
            minScore = inputArr[i]
            countMin ++
        }
        if inputArr[i] > maxScore {
            maxScore = inputArr[i]
            countMax ++
        }
    }
    
    minMax = append(minMax, countMax)
    minMax = append(minMax, countMin)
    return minMax
}

func main() {
    var scores []int32
    scores = []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
    result := breakingRecords(scores)
    fmt.Println(result)
}

