// Tower of Hanoi in Golang
package main
  
import "fmt"
 
type solver interface {
    play(int)
}
   
// towers is example of type satisfying solver interface
type towers struct {
    // an empty struct
}
  
// play is sole method required to implement solver type
func (t *towers) play(n int) {    
    t.moveN(n, 3, 4, 5)
}
  
// recursive algorithm
func (t *towers) moveN(n, from, to, via int) {
    if n > 0 {
        t.moveN(n-1, from, via, to)
        t.moveM(from, to)
        t.moveN(n-1, via, to, from)
    }
}
 
func (t *towers) moveM(from, to int) {
    fmt.Println("Move disk from rod", from, "to rod", to)
}
 
func main() {
    var t solver    
    t = new(towers) // type towers must satisfy solver interface
    t.play(4)
    } 
Move disk from rod 3 to rod 5
Move disk from rod 3 to rod 4
Move disk from rod 5 to rod 4
Move disk from rod 3 to rod 5
Move disk from rod 4 to rod 3
Move disk from rod 4 to rod 5
Move disk from rod 3 to rod 5
Move disk from rod 3 to rod 4
Move disk from rod 5 to rod 4
Move disk from rod 5 to rod 3
Move disk from rod 4 to rod 3
Move disk from rod 5 to rod 4
Move disk from rod 3 to rod 5
Move disk from rod 3 to rod 4
Move disk from rod 5 to rod 4

Program exited.