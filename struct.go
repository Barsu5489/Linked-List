package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (ll *LinkedList) AddNode(data int) {
	newNode := &Node{data: data, next: ll.head}
	ll.head = newNode
}
func (ll *LinkedList) PrintLength() {
	cur := ll.head
	count := 0
	for cur != nil {
		count++
		cur = cur.next
	}
	fmt.Println(count)
}
func (ll *LinkedList) PrintList() {
	cur := ll.head
	for cur != nil {
		fmt.Print("->", cur.data)
		cur = cur.next
	}
	fmt.Println()
}

func (ll *LinkedList) DeleteNode(n int) {

	if ll.head != nil && ll.head.data == n {
		ll.head = ll.head.next
		return
	}


	current := ll.head
	for current != nil && current.next != nil {
		if current.next.data == n {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func main() {
	myLinkedList := LinkedList{}
	myLinkedList.AddNode(3)
	myLinkedList.AddNode(5)
	myLinkedList.AddNode(6)
	myLinkedList.AddNode(8)
	myLinkedList.DeleteNode(5)
	myLinkedList.PrintList()
	myLinkedList.PrintLength()

}
