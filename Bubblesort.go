//1.bubble sort
package main

import (
    "fmt"
)

var toBeSorted [10]int = [10]int{1,3,2,4,8,6,7,2,3,0}

func bubbleSort(input [10]int) {
    // n is the number of items in our list
    n := 10
    // set swapped to true
    swapped := true
    // loop
    for swapped {
        // set swapped to false
        swapped = false
        // iterate through all of the elements in our list
        for i := 1; i < n; i++ {
            // if the current element is greater than the next
            // element, swap them
            if input[i-1] > input[i] {
                // log that we are swapping values for posterity
                fmt.Println("Swapping")
                // swap values using Go's tuple assignment
                input[i], input[i-1] = input[i-1], input[i]
                // set swapped to true - this is important
                // if the loop ends and swapped is still equal
                // to false, our algorithm will assume the list is
                // fully sorted.
                swapped = true
            }
        }
    }
    // finally, print out the sorted list
    fmt.Println(input)
}

func main() {
    fmt.Println("this is bubble sort")
    bubbleSort(toBeSorted)
}
//3.bubble sort
package main

import "fmt"

func main() {
    sample := []int{3, 4, 5, 2, 1,8,9,0,6,-9,90}
    bubbleSort(sample)
    sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
    bubbleSort(sample)
}

func bubbleSort(arr []int) {
    len := len(arr)
    for i := 0; i < len-1; i++ {
        for j := 0; j < len-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    fmt.Println("\nAfter Bubble Sorting")
    for _, val := range arr {
        fmt.Println(val)
    }
}
After Bubble Sorting
-9
0
1
2
3
4
5
6
8
9
90

After Bubble Sorting
-3
-1
1
2
3
4
5
7
8