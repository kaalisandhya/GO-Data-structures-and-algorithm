1.package main
 
import (
	"fmt"
	"math/rand"
)
type Node struct {
	num  int
	prev *Node
	next *Node
}
 
type List struct {
	tail  *Node
	start *Node
}
func main() {
	items := &List{}
	size := 10
	//rand_number := make([]int, size, size)
	for i := 0; i < size; i++ {
		node := Node{num: rand.Intn(100)}
		if node.num == 0 {
			i = i - 1
			continue
 
		}
		items.insertNode(&node)
		fmt.Printf("%v and number is %v\n", i, node.num)
	}
	items.Display()
	items.DisplayReverse()
}
 
func (l *List) Display() {
	list := l.start
	for list != nil {
		fmt.Printf("value = %v and prev = %v and next= %v\n", list.num, list.prev, list.next)
		list = list.next
	}
	fmt.Println()
}
 
func (l *List) DisplayReverse() {
	list := l.tail
	for list != nil {
		fmt.Printf("value = %v\n", list.num)
		list = list.prev
	}
	fmt.Println()
}
 
func (l *List) insertNode(newNode *Node) {
	if l.start == nil {
		l.start = newNode
		l.tail = newNode
	} else {
		currentNode := l.start
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		l.tail = newNode
	}
}
0 and number is 81
1 and number is 87
2 and number is 47
3 and number is 59
4 and number is 81
5 and number is 18
6 and number is 25
7 and number is 40
8 and number is 56
9 and number is 94
value = 81 and prev = <nil> and next= &{87 0xc0000ae040 0xc0000ae080}
value = 87 and prev = &{81 <nil> 0xc0000ae060} and next= &{47 0xc0000ae060 0xc0000ae0a0}
value = 47 and prev = &{87 0xc0000ae040 0xc0000ae080} and next= &{59 0xc0000ae080 0xc0000ae0c0}
value = 59 and prev = &{47 0xc0000ae060 0xc0000ae0a0} and next= &{81 0xc0000ae0a0 0xc0000ae0e0}
value = 81 and prev = &{59 0xc0000ae080 0xc0000ae0c0} and next= &{18 0xc0000ae0c0 0xc0000ae100}
value = 18 and prev = &{81 0xc0000ae0a0 0xc0000ae0e0} and next= &{25 0xc0000ae0e0 0xc0000ae120}
value = 25 and prev = &{18 0xc0000ae0c0 0xc0000ae100} and next= &{40 0xc0000ae100 0xc0000ae140}
value = 40 and prev = &{25 0xc0000ae0e0 0xc0000ae120} and next= &{56 0xc0000ae120 0xc0000ae180}
value = 56 and prev = &{40 0xc0000ae100 0xc0000ae140} and next= &{94 0xc0000ae140 <nil>}