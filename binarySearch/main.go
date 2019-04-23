package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var t Tree
	t.Initial()
	fmt.Println("initialed t:",t)
	//生成10个节点并加入二叉树
	for i := 0; i < 10; i ++ {
		//key 值为0-50之间的随机整数数
		key := rand.Intn(50)
		//fmt.Println("key:", key)

		var i Node
		i.key = key
		//定义value值为对应key值的10倍
		i.value = 10*key
		t.Add(&i)
	}
	/*
	var list List
	var n ListNode
	n.Value = 1

	list.Initial()
	list.Add(&n)
	fmt.Println("head",list.head)
	*/
	t.LevelOrder()
}
