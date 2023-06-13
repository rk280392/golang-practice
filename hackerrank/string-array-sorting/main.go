package main

//TO DO

func RemainderSorting(strArr []string) {

	for i, j := 0, 0; i < len(strArr); i++ {
		lenI := len(strArr[i])
		lenJ := len(strArr[j])
		modI := lenI % 3
		modJ := lenJ % 3

	}
}

func main() {
	strArr := []string{"Rajesh", "kumar", "hero"}
	RemainderSorting(strArr)
}
