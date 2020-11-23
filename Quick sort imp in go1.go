1.package main

import (
	"fmt"
	"strconv"
)

func quickSort(arr []int, low, high int) {
	if low < high {
		var pivot = partition(arr, low, high)
		quickSort(arr, low, pivot)
		quickSort(arr, pivot + 1, high)
	}
}

func partition(arr []int, low, high int) int {
	var pivot = arr[low]
	var i = low
	var j = high

	for i < j {
		for arr[i] <= pivot && i < high {
			i++;
		}
		for arr[j] > pivot && j > low {
			j--
		}

		if i < j {
			var temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	arr[low] = arr[j]
	arr[j] = pivot

	return j
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(strconv.Itoa(arr[i]) + " ")
	}
	fmt.Println("")
}

func main() {
	var arr = []int { 18, 3, 11, 6, -9, 1, 0 , 5 , 90 , 56 , 7 , 8 }

	fmt.Print("Before Sorting: ")
	printArray(arr)

	quickSort(arr, 0, len(arr) - 1)
	fmt.Print("After Sorting: ")
	printArray(arr)
}
Before Sorting: 18 3 11 6 -9 1 0 5 90 56 7 8 
After Sorting: -9 0 1 3 5 6 7 8 11 18 56 90 
Program exited.