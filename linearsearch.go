1.// Linear Search in Golang
package main
  
import "fmt"
 
func linearsearch(datalist []int, key int) bool {
    for _, item := range datalist {
        if item == key {
            return true
        }
    }
    return false
} 
  
func main() {
    items := []int{95,78,46,58,45,86,99,251,320,89,56,78}
    fmt.Println(linearsearch(items,58))
}
true

Program exited.
2.//linear search in golang
package main

import (
	"fmt"
)

func linearSearch(array []int, size int, num int) {
	for i := 0; i < size; i++ {
		if array[i] == num {
			fmt.Println("Number found at position : ", i+1)
			return
		}
	}

	fmt.Println("Number not found")
	return

}

func main() {
	fmt.Println("Enter the size of array")
	var size int
	fmt.Scan(&size)

	var num int
	array := make([]int, size, 100)

	for i := 0; i < size; i++ {
		fmt.Println("Enter element : ")
		fmt.Scan(&array[i])
	}

	fmt.Println("Enter the number to be searched")
	fmt.Scan(&num)

	linearSearch(array, size, num)

}
    return true
}
 
 
func main(){
    items := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
    fmt.Println(binarySearch(100, items))
}