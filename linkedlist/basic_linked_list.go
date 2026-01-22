package linkedlist

import "fmt"

type Node struct {
	val  int
	next *Node
}

func NewNode(val int) *Node {
	return &Node{
		val:  val,
		next: nil,
	}
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
	}
}
func (this *LinkedList) insertBeginning(val int) {
	newNode := NewNode(val)
	newNode.next = this.head
	this.head = newNode
}

func (this *LinkedList) ListVal() {
	temp := this.head
	if temp == nil {
		fmt.Println("List is empty")
		return
	}
	for {
		fmt.Println(temp.val)
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
}

func (this *LinkedList) insertEnd(val int) {
	newNode := NewNode(val)
	temp := this.head
	if temp == nil {
		this.head = newNode
		this.tail = newNode
		return
	}
	// main operation is from here
	fmt.Println(this.tail)
}
func (this *LinkedList) deleteByVal(val int) {
	curr := this.head
	if curr == nil {
		fmt.Println("list is empty!")
		return
	}

}
func (this *LinkedList) deleteByEnd(val int) {

}
func (this *LinkedList) deleteBeginning(val int) {
	if this.head == nil {
		fmt.Println("list is empty!")
		return
	}
	this.head = this.head.next
}
func BasicLinkedList() {
	list := NewLinkedList()
	list.insertBeginning(3)
	list.insertBeginning(2)
	list.insertBeginning(1)
	list.insertEnd(4)
	list.ListVal()
}
