package main

import "fmt"

type ListObject interface{
}
type ListNode struct {
	Value ListObject
	Next *ListNode
}
type List struct {
	size int
	head *ListNode
	tail *ListNode
}
func (l *List) Initial() {
	l.size = 0
	l.head = nil
	l.tail = nil
}
func (l *List) Add(node *ListNode) bool {
	// fmt.Println(node)
	//  fmt.Println(l.size)
	if l.size == 0 {
		l.head = node
		l.tail = node
		l.size ++
	} else {

		oldTail := l.tail
		//fmt.Println("oldTail:",oldTail)

		oldTail.Next = node
		//fmt.Println("oldTail.Next",oldTail.Next)
		l.tail = node
		l.size ++

	}
	return true
}

func (l *List) InsertToFirst(node *ListNode) bool {
	if node == nil {
		return false
	} else {
		l.head = node
		node.Next = l.head
		l.size ++
	}
	return true
}
func (l *List)  Find(i int) *ListNode {
	if i > l.size {
		return nil
	}
	//fmt.Println(i)
	item := l.head
	for j := 0; j < i; j ++ {
		item = item.Next
	}
	return item

}
func (l *List) delete(i int) bool{
	if i < 1 || i > l.size {
		return false
	}
	if i == 1 {
		l.head = l.head.Next
	} else {
		m := l.head
		for k := 0; k < i -2; k ++ {

			m = m.Next
			fmt.Println("m:",m)
		}
		cur := m
		fmt.Println("cur m", cur)
		m.Next = cur.Next.Next
		fmt.Println("mmm",m.Next)

	}
	l.size --



	return true

}
func (l *List) Push(node interface{})  {
	var listNode ListNode
	listNode.Value = node
	l.Add(&listNode)
}
func (l *List) Pop() interface{} {
	result := l.head.Value
	l.delete(1)

	return result
}
func (l *List) Print() {
	cur := l.head
	//fmt.Print(cur)
	//fmt.Print(l.size)

	for i := 0; i < l.size; i ++ {
		fmt.Println(cur)
		cur = cur.Next

	}
}

func main() {
	var NewList List
	//NewList.Initial()
	fmt.Println(NewList)

	for i := 1; i <= 10; i ++ {
		var node ListNode
		node.Value = i
		node.Next = nil
		NewList.Add(&node)
	}


	for j := 0; j < 10; j ++ {
		var result *ListNode
		result = NewList.Find(j)
		fmt.Println("value:", result)
	}
	fmt.Print("size",NewList.size)
	//fmt.Println("1111111")
	NewList.delete(1)
	fmt.Print("new size",NewList.size)
	//fmt.Println(NewList.head)
	/*
	fmt.Println(NewList.Find(2))
	fmt.Println(NewList.head)
	*/
	/*
	for j := 0; j < 10; j ++ {
		var result *ListNode
		result = NewList.Find(j)
		fmt.Println("nvalue:", result)
	}
	fmt.Println("head",NewList.head)
	*/

}
