//sorting
package main

	import "fmt"
import "sort"

	func main() {
// Sort methods are specific to the builtin type; here’s an example for strings. Note that sorting is in-place, so it changes the given slice and doesn’t return a new one.
strs := []string{"a", "i", "g","d","l","f","m","h","o"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)


	
	//An example of sorting ints.
ints := []int{7, 3, 5,8,0,2,9,10,4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

/// We can also use sort to check if a slice is already in sorted order.
	    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}
Strings: [a d f g h i l m o]
Ints:    [0 2 3 4 5 7 8 9 10]
Sorted:  true
