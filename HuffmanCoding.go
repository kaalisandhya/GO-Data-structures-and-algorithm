// Huffman Coding in Golang
package main
  
import (
    "container/heap"
    "fmt"
)
  
type HuffmanTree interface {
    Freq() int
}
  
type HuffmanLeaf struct {
    freq  int
    value rune
}
  
type HuffmanNode struct {
    freq        int
    left, right HuffmanTree
}
  
func (self HuffmanLeaf) Freq() int {
    return self.freq
}
  
func (self HuffmanNode) Freq() int {
    return self.freq
}
  
type treeHeap []HuffmanTree
  
func (th treeHeap) Len() int { return len(th) }
func (th treeHeap) Less(i, j int) bool {
    return th[i].Freq() < th[j].Freq()
}
func (th *treeHeap) Push(ele interface{}) {
    *th = append(*th, ele.(HuffmanTree))
}
func (th *treeHeap) Pop() (popped interface{}) {
    popped = (*th)[len(*th)-1]
    *th = (*th)[:len(*th)-1]
    return
}
func (th treeHeap) Swap(i, j int) { th[i], th[j] = th[j], th[i] }
 
// The main function that builds a Huffman Tree and print codes by traversing
// the built Huffman Tree
func buildTree(symFreqs map[rune]int) HuffmanTree {
    var trees treeHeap
    for c, f := range symFreqs {
        trees = append(trees, HuffmanLeaf{f, c})
    }
    heap.Init(&trees)
    for trees.Len() > 1 {
        // two trees with least frequency
        a := heap.Pop(&trees).(HuffmanTree)
        b := heap.Pop(&trees).(HuffmanTree)
  
        // put into new node and re-insert into queue
        heap.Push(&trees, HuffmanNode{a.Freq() + b.Freq(), a, b})
    }
    return heap.Pop(&trees).(HuffmanTree)
}
 
// Prints huffman codes from the root of Huffman Tree.  It uses byte[] to
// store codes
func printCodes(tree HuffmanTree, prefix []byte) {
    switch i := tree.(type) {
    case HuffmanLeaf:
        // If this is a leaf node, then it contains one of the input
        // characters, print the character and its code from byte[]
        fmt.Printf("%c\t%d\t%s\n", i.value, i.freq, string(prefix))
    case HuffmanNode:
        // Assign 0 to left edge and recur
        prefix = append(prefix, '0')
        printCodes(i.left, prefix)
        prefix = prefix[:len(prefix)-1]
  
        // Assign 1 to right edge and recur
        prefix = append(prefix, '1')
        printCodes(i.right, prefix)
        prefix = prefix[:len(prefix)-1]
    }
}
 
// Driver program to test above functions
func main() {
    test := "abcdefghijklmnopqrstuvwxyz"
  
    symFreqs := make(map[rune]int)
    // read each symbol and record the frequencies
    for _, c := range test {
        symFreqs[c]++
    }
  
    // example tree
    exampleTree := buildTree(symFreqs)
  
    // print out results
    fmt.Println("SYMBOL\tWEIGHT\tHUFFMAN CODE")
    printCodes(exampleTree, []byte{})
}
SYMBOL	WEIGHT	HUFFMAN CODE
q	1	0000
m	1	0001
t	1	0010
f	1	0011
e	1	0100
v	1	0101
c	1	01100
x	1	01101
z	1	01110
r	1	01111
n	1	10000
p	1	10001
k	1	10010
a	1	10011
i	1	10100
s	1	10101
g	1	10110
b	1	10111
d	1	11000
u	1	11001
j	1	11010
o	1	11011
y	1	11100
l	1	11101
w	1	11110
h	1	11111
