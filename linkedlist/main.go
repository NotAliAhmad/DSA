package main

import "fmt"

// Node represents a node in the linked list.
type Node struct {
    data int
    next *Node
}

// LinkedList represents a singly linked list.
type LinkedList struct {
    head *Node
}

// Append adds a new node with the given data to the end of the linked list.
func (ll *LinkedList) Append(data int) {
    newNode := &Node{data: data, next: nil}

    if ll.head == nil {
        ll.head = newNode
        return
    }

    current := ll.head
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}

// Display prints the elements of the linked list.
func (ll *LinkedList) Display() {
    current := ll.head
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}

func main() {
    // Creating a linked list
    linkedList := LinkedList{}

    // Appending elements to the linked list
    linkedList.Append(1)
    linkedList.Append(2)
    linkedList.Append(3)

    // Displaying the linked list
    linkedList.Display()
}
