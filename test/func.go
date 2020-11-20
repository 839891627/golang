package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func printArray(arr []int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	arr[0] = 100
}
func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	//var arr1 []int
	//arr2:=[3]int{1,2,3}
	arr3 := []int{1, 2, 3, 4, 5}
	//printArray(arr1)
	//printArray(arr2)
	printArray(arr3)
	fmt.Println(arr3)
}
