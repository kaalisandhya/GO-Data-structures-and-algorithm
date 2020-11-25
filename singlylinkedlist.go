1.package main

import "fmt"

type ele struct {
    name string
    next *ele
}

type singleList struct {
    len  int
    head *ele
}

func initList() *singleList {
    return &singleList{}
}

func (s *singleList) AddFront(name string) {
    ele := &ele{
        name: name,
    }
    if s.head == nil {
        s.head = ele
    } else {
        ele.next = s.head
        s.head = ele
    }
    s.len++
    return
}

func (s *singleList) AddBack(name string) {
    ele := &ele{
        name: name,
    }
    if s.head == nil {
        s.head = ele
    } else {
        current := s.head
        for current.next != nil {
            current = current.next
        }
        current.next = ele
    }
    s.len++
    return
}

func (s *singleList) RemoveFront() error {
    if s.head == nil {
        return fmt.Errorf("List is empty")
    }
    s.head = s.head.next
    s.len--
    return nil
}

func (s *singleList) RemoveBack() error {
    if s.head == nil {
        return fmt.Errorf("removeBack: List is empty")
    }
    var prev *ele
    current := s.head
    for current.next != nil {
        prev = current
        current = current.next
    }
    if prev != nil {
        prev.next = nil
    } else {
        s.head = nil
    }
    s.len--
    return nil
}

func (s *singleList) Front() (string, error) {
    if s.head == nil {
        return "", fmt.Errorf("Single List is empty")
    }
    return s.head.name, nil
}
AddFront: A
AddFront: B
AddBack: C
Size: 3
B
A
C
RemoveFront
RemoveBack
RemoveBack
RemoveBack
RemoveBack Error: removeBack: List is empty
TranverseError: List is empty
Size: 0
Program exited.
2.package main

import "fmt"

// Node can store two values, 'id' and 'name'.
// Another value, 'ptr' is a pointer to another node
type Node struct {
	id   int
	name string
	ptr  *Node
}

// LinkedList struct
type LinkedList struct {
	head *Node
}

// LinkedList method to append a node to tail
func (linkedlist *LinkedList) append(newnode *Node) *LinkedList {
	if linkedlist.head == nil {
		linkedlist.head = newnode
		newnode.ptr = nil
		return linkedlist
	} // else
	// initialization; condition; increment
	for n := linkedlist.head; n != nil; n = n.ptr {
		if n.ptr == nil {
			n.ptr = newnode
			return linkedlist
		}
	}
	return linkedlist
}

// LinkedList method to print the nodes details
func (linkedlist *LinkedList) print() {
	for n := linkedlist.head; n != nil; n = n.ptr {
		fmt.Println(n.id, n.name, n.ptr)
	}
}

// LinkedList method to insert a new node before a particular node
func (linkedlist *LinkedList) insertBefore(name string, newNode *Node) bool {
	if linkedlist.head == nil {
		return false
	}

	ptrNode := linkedlist.head // points to first node
	for ptrNode.ptr != nil {   // while() loop
		if ptrNode.ptr.name == name {
			newNode.ptr = ptrNode.ptr // new node points to the next node
			ptrNode.ptr = newNode     // current node points to the new node
			return true
		}
		ptrNode = ptrNode.ptr // points to next node
	}
	return false
}

func main() {
	// create a linked list object
	linkedlist := &LinkedList{}

	// create nodes
	nodeAbc := &Node{1, "Abc", nil}
	nodeBac := &Node{2, "Bac", nil}
	nodeCab := &Node{3, "Cab", nil}
	nodeDab := &Node{4, "Dab", nil}
	nodeEac := &Node{5, "Eac", nil}

	// append the nodes to the linked list
	linkedlist.append(nodeAbc)
	linkedlist.append(nodeBac)
	linkedlist.append(nodeCab)
	linkedlist.append(nodeDab)
	linkedlist.append(nodeEac)

	// print the linked list
	linkedlist.print()

	// create a new node
	nodeFab := &Node{6, "Fab", nil}
	// insert the new node before node "Bac"
	linkedlist.insertBefore("Bac", nodeFab)
	// print the linked list
	linkedlist.print()
}
	
1 Abc &{2 Bac 0xc00000c0e0}
2 Bac &{3 Cab 0xc00000c100}
3 Cab &{4 Dab 0xc00000c120}
4 Dab &{5 Eac <nil>}
5 Eac <nil>
1 Abc &{6 Fab 0xc00000c0c0}
6 Fab &{2 Bac 0xc00000c0e0}
2 Bac &{3 Cab 0xc00000c100}
3 Cab &{4 Dab 0xc00000c120}
4 Dab &{5 Eac <nil>}
5 Eac <nil>

Program exited
