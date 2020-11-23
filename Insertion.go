//3.insertion sort
package main

import "fmt"

func main() {
    sample := []int{3, 4, 5, 2, 1,9,0,-9,-6,7}
    insertionSort(sample)

    sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
    insertionSort(sample)
}

func insertionSort(arr []int) {
    len := len(arr)
    for i := 1; i < len; i++ {
        for j := 0; j < i; j++ {
            if arr[j] > arr[i] {
                arr[j], arr[i] = arr[i], arr[j]
            }
        }
    }
    
    fmt.Println("After Sorting")
    for _, val := range arr {
        fmt.Println(val)
    }
}
After Sorting
-9
-6
0
1
2
3
4
5
7
9
After Sorting
-3
-1
1
2
3
4
5
7
8