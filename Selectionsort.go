//4.selection sort
package main

import "fmt"

func main() {
    sample := []int{33, 44, 55, 22, 11,10,99,88,66}
    selectionSort(sample)
    sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3,90,78,567,89}
    selectionSort(sample)
}

func selectionSort(arr []int) {
    len := len(arr)
    for i := 0; i < len-1; i++ {
        minIndex := i
        for j := i + 1; j < len; j++ {
            if arr[j] < arr[minIndex] {
                arr[j], arr[minIndex] = arr[minIndex], arr[j]
            }
        }
    }
    fmt.Println("\nAfter SelectionSort")
    for _, val := range arr {
        fmt.Println(val)
    }
}
After SelectionSort
10
11
22
33
44
55
66
88
99

After SelectionSort
-3
-1
1
2
3
4
5
7
8
78
89
90
567