package main

import "fmt"

//     |4.2|3.3|4.nil|

//  	move to node.data == 3
//      node.data == 0 and node.next == nil
// 		node.next == nextNode

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Insert(data int) {

	newNode := &Node{
		data: data,
	}

	currentNode := ll.head

	if currentNode == nil {
		ll.head = newNode
	} else {
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}

}

func (ll *LinkedList) Display() {

	if ll.head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	fmt.Printf("LinkedList >>")

	currentNode := ll.head

	for currentNode != nil {
		fmt.Printf(" : %v ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()

}

func (ll *LinkedList) Remove(data int) {

}

func main() {

	newLinkedList := &LinkedList{}

	newLinkedList.Insert(2)
	newLinkedList.Insert(8)
	newLinkedList.Insert(3)
	newLinkedList.Insert(6)
	newLinkedList.Display()

}
