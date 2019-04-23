package main
import "fmt"
type Object interface{
}
type Node struct {
	Value Object
	Next *Node
}
type List struct {
	size int
	head *Node
	tail *Node
}
func (l *List) Initial() {
	l.size = 0
	l.head = nil
	l.tail = nil
}
func (l *List) Add(node *Node) bool {
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
func (l *List) InsertToFirst(node *Node) bool {
	if node == nil {
		return false
	} else {
		l.head = node
		node.Next = l.head
		l.size ++
	}
	return true
}
func (l *List)  Find(i int) *Node {
	if i > l.size {
		return nil
	}
	fmt.Println(i)
	item := l.head
	for j := 0; j < i; j ++ {
		item = item.Next
	}
	return item

}
func (l *List) Push()  {

}
func main() {
	var NewList List
	//NewList.Initial()
	fmt.Println(NewList)

	for i := 1; i <= 10; i ++ {
		var node Node
		node.Value = i
		node.Next = nil
		NewList.Add(&node)
	}

	for j := 0; j < 10; j ++ {
		var result *Node
		result = NewList.Find(j)
		fmt.Println("value:", result)
	}
	/*
	fmt.Println(NewList.Find(2))
	fmt.Println(NewList.head)
	*/
}