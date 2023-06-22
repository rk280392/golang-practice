package main

import "fmt"

func beautifulTriplets(d int32, arr []int32) int32 {
	var count int32 = 0
	//n := len(arr)
	frequencyMap := make(map[int32]int32)

	for _, num := range arr {
		frequencyMap[num]++
	}

	fmt.Println(frequencyMap)

	for i := range arr {
		if frequencyMap[arr[i]+d] > 0 && frequencyMap[arr[i]+2*d] > 0 {
			count++
			fmt.Printf("%d %d %d\n", arr[i], arr[i]+d, arr[i]+2*d)
		}
	}
	return count
}

func main() {

	var d int32 = 3
	arr := []int32{1, 2, 4, 5, 7, 8, 10}

	/*

		Take input as stdin instead of passing in the main function

			var n, d int32
			fmt.Scan(&n, &d)
			arr := make([]int32, n)
				for i := 0; i < int(n); i++ {
					fmt.Scan(&arr[i])
				}
	*/
	result := beautifulTriplets(d, arr)
	fmt.Println(result)
}

/*

Explanation:

	a[j] - a[i] = d --> a[j] = d + a[i]

	a[k] - a[j] = d --> a[k] = d + a[j] --> a[k] = d + d + a[i] --> a[k] = 2d + a[i]

	a[i] will be the loop i need to find second and third element. Check if a[i] + d and a[i] + 2d exists in the frequence map. If yes, this will be a triplet.

*/
