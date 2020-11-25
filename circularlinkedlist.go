1.package main
 
import (
   "container/ring"
   "fmt"
)
 
func main() {
   numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12}
   r := ring.New(len(numbers))
 
   // initialize ring
   for i := 0; i < r.Len(); i++ {
      r.Value = numbers[i]
      r = r.Next()
   }
 
   //Do calls function f on each element of the ring, in forward order.
   r.Do(func(x interface{}) {
      fmt.Print(x)
   })
   fmt.Println()
 
   // Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0)
   // in the ring and returns that ring element. r must not be empty.
   //r = r.Move(5)
 
   // print elements in reverse order
   for i := 0; i < r.Len(); i++ {
      fmt.Print(r.Value)
      r = r.Prev()
   }
   fmt.Println()
 
   // Move moves n % r.Len() elements backward (n < 0) or forward (n >= 0)
   // in the ring and returns that ring element. r must not be empty.
   r = r.Move(5)
   r.Do(func(x interface{}) {
      fmt.Print(x)
   })
   fmt.Println()  
}
123456789101112
112111098765432
678910111212345

Program exited.