package main

import "fmt"

func main(){
	//arrays
	var arr [10]int

	for i:=0; i<len((arr)); i++ {
		arr[i] = i + 1
	}
	fmt.Println("Array:", arr)

	arr1 := [3]int{1,2,3}
	fmt.Println("Array1:", arr1)

	//slices
	sli := []int{1,2,3,4,5}
	fmt.Println("Slice:", sli)
	sli = append(sli, 6,7,8)
	fmt.Println("Slice after append:", sli)

	sli1 := make([]int, 5)
	for i := 0; i < len(sli1); i++ {
		sli1[i] = i + 1
	}
	fmt.Println("Slice1:", sli1)
}