1.//Declaring an Integer or String Array
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var intArray [5]int
	var strArray [5]string

	fmt.Println(reflect.ValueOf(intArray).Kind())
	fmt.Println(reflect.ValueOf(strArray).Kind())
}
array
array

Program exited.
2.//assign and access array element values in Go
package main

import "fmt"

func main() {
	var theArray [3]string
	theArray[0] = "kodur"  
	theArray[1] = "kadapa" 
	theArray[2] = "siripiuram" 

	fmt.Println(theArray[0]) 
	fmt.Println(theArray[1]) 
	fmt.Println(theArray[2]) 
}
kodur
kadapa
siripiuram

Program exited.
3.//initialize an Array with an Array Literal in Go
package main

import "fmt"

func main() {
	x := [5]int{10, 20, 30, 40, 50}   
	var y [5]int = [5]int{10, 20, 30} 

	fmt.Println(x)
	fmt.Println(y)
}
[10 20 30 40 50]
[10 20 30 0 0]

Program exited.

4.package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := [...]int{10, 20, 30}

	fmt.Println(reflect.ValueOf(x).Kind())
	fmt.Println(len(x))
}
array
3

Program exited.
//5.iterate over an Array using for loop
package main

import "fmt"

func main() {
	intArray := [5]int{10, 20, 30, 40, 50}

	
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}

	
	for index, element := range intArray {
		fmt.Println(index, "=>", element)

	}

	
	for _, value := range intArray {
		fmt.Println(value)
	}

	j := 0
	
	for range intArray {
		fmt.Println(intArray[j])
		j++
	}
}
10
20
30
40
50
0 => 10
1 => 20
2 => 30
3 => 40
4 => 50
10
20
30
40
50
10
20
30
40
50
//6.Copy an array by value and reference into another array
package main

import "fmt"

func main() {

	strArray1 := [3]string{"kadapa", "kodur", "peeleru"}
	strArray2 := strArray1  
        strArray3 := &strArray1

	fmt.Printf("strArray1: %v\n", strArray1)
        fmt.Printf("strArray2: %v\n", strArray2)	

        strArray1[0] = "kodur"
  
	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)	
	fmt.Printf("*strArray3: %v\n", *strArray3)	
}
strArray1: [kadapa kodur peeleru]
strArray2: [kadapa kodur peeleru]
strArray1: [kodur kodur peeleru]
strArray2: [kadapa kodur peeleru]
*strArray3: [kodur kodur peeleru]

Program exited.