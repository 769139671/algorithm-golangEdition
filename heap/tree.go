package main

import "fmt"

type Obj interface {

}
type Node1 struct {
	Value Obj
	Left *Node1
	Right *Node1
}
func createNode(value int) *Node1 {
	return &Node1{value,nil,nil}
}
//LNR
func preRange(node *Node1)  {
	if node == nil {
		return
	}
	fmt.Println(node)
	preRange(node.Left)
	//if node.left == nil
	//print node

	preRange(node.Right)
}
/*
func Add(v int) bool {
	value := createNode(v)

}
/*
func main() {

	root := Node1{3,nil,nil}
	//fmt.Println(root)zz
	root.Left = createNode(0)
	root.Right = createNode(5)
	root.Left.Right = createNode(4)
	root.Right.Left = createNode(8)
	/*
	fmt.Println(root.Left)
	fmt.Println(root.Left.Right)
	fmt.Println(root.Right)
	*/
	//preRange(&root)
