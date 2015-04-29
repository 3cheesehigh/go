package main

import "fmt"

//a)
//TODO extract as DoubleLikedList package

type Link struct {
	Cont any
	next,
	prev *Link
}

type doubleLinkedList struct {
	head   Link
	length int
}

type any interface{}

//b)
func (list *doubleLinkedList) insEle(n int, newEle *Link) {
	cur := &list.head
	if n > list.length {
		print("Index out of List")
		return
	}
	for x := 0; x < n; x++ {
		cur = cur.next
	}
	tmp := cur.next
	tmp.prev = newEle
	newEle.prev = cur
	newEle.next = tmp
	cur.next = newEle

	list.length++
}

func (list *doubleLinkedList) delEle(n int) {
	if n > list.length {
		print("Index out of List")
		return
	}
	cur := list.head.next
	for x := 0; x < n; x++ {
		cur = cur.next
	}
	cur.next.prev = cur.prev
	cur.prev.next = cur.next
	list.length--

}

func (list *doubleLinkedList) PrintDoubleLinkedList() {
	cur := list.head.next
	for i := 0; i < list.length; i++ {
		fmt.Printf("%d ", cur.Cont)
		cur = cur.next

	}
	fmt.Printf("\n")
}

func main() {
	//New list with sentinal
	list := new(doubleLinkedList)
	list.head.next = &list.head
	list.head.prev = &list.head
	newEle1 := Link{Cont: 5}
	newEle2 := Link{Cont: 6}
	newEle3 := Link{Cont: 7}

	list.insEle(0, &newEle1)
	list.insEle(1, &newEle2)
	list.insEle(2, &newEle3)

	fmt.Printf("Listenlänge vorm Löschen: %d\n", list.length)
	list.PrintDoubleLinkedList()

	list.delEle(1)
	fmt.Printf("Listenlänge nach dem  Löschen: %d\n", list.length)
	list.PrintDoubleLinkedList()
}
