package main

import (
	"fmt"
	"math/rand"
)
type Object interface {
}
type Node struct {
	key int
	value Object
	//parent *Node
	left *Node
	right *Node
}
type Tree struct {
	root *Node
	count int
}

func (t *Tree) Initial() {
	t.count = 0
}

func (t *Tree) Count() {
	t.count ++
}
func (t *Tree)Add(node *Node) {
	if t.count == 0 {
		t.root = node
		t.Count()
	} else {
		t.insert(node, t.root)
	}
}
func (t *Tree)insert(node *Node, current *Node) {
	if node.key == current.key {
		current.value = node.value
		return
	}
	if node.key > current.key {
		if current.right == nil {
			current.right = node
			t.Count()
			return
		} else {
			current = current.right
			t.insert(node,current)
		}
	}
	if node.key < current.key {
		if current.left == nil {
			current.left = node
			t.Count()
			return
		} else {
			current = current.left
			t.insert(node, current)
		}
	}
}

func (t *Tree) Search(k int) interface{}{
	a := t.Contain(k)
	if a == false {
		fmt.Println("not find")
		return nil
	} else {
		return t.search(k,t.root)
	}
}
func (t *Tree) search(k int, current *Node) interface{} {
	if k == current.key {
		return current.value
	}
	if k > current.key {
		current = current.right
		return t.search(k, current)
	}
	if k < current.key {
		current = current.left
		return t.search(k, current)
	}
	return nil
}
func (t *Tree) Contain(k int) bool {
	return t.contain(t.root, k)
}
func (t *Tree) contain(current *Node,k int) bool{
	if current.key == k {
		return true
	}
	if k > current.key {

		if current.right == nil {
			return false
		} else {
			current = current.right
			return t.contain(current,k)
		}
	}
	if k < current.key {
		if current.left == nil {
			return false
		} else {
			current = current.left
			return t.contain(current, k)
		}
	}

	fmt.Println("f")
	return false
}
func (t *Tree) PreOrder() {
	//fmt.Printf("key:%v , value:%v\n",t.root.key, t.root.value)
	t.preOrder(t.root)

}
func (t *Tree) preOrder(current *Node)  {
	fmt.Printf(
		"key:%v , value:%v, left: %v, right:%v\n",current.key, current.value, current.left, current.right)

	if current.left != nil {
		t.preOrder(current.left)
	}

	if current.right != nil {
		t.preOrder(current.right)
	}
}
/*
func (t *Tree) LevelOrder() {
	var list List

	t.levelOrder(t.root,&list)
	fmt.Print("start print:\n")
	list.Print()
	return
}
func (t *Tree) levelOrder(cur *Node,list *List)  {
	list.Push(cur)
	if cur.left != nil || cur.right != nil {
		if cur.left != nil {
			cur = cur.left
			t.levelOrder(cur, list)
		} else {
			cur = cur.right
			t.levelOrder(cur, list)
		}
	} else {

		return
	}
	/*
	if cur.left != nil {
		cur := cur.left
		t.levelOrder(cur, list)
	}
	if cur.right != nil {
		cur = cur.right
		t.levelOrder(cur, list)
	}


}
*/
func (t *Tree) LevelOfChan() {
	var c chan *Node
	c = make(chan *Node,100)
	t.levelOfChan(t.root,c)
}
func (t *Tree) levelOfChan(node *Node,c chan *Node) {
	t.push(node,c)
	for i := 0; i <t.count; i ++ {
		t.pop(c)
	}
}
func (t *Tree) pop(c chan *Node) {
	//p := <-c
	for i := 0; i < len(c); i ++ {
		p := <- c
		fmt.Println("p:",p)
		if p.left != nil && p.right != nil {
			//fmt.Println("111")
			t.push(p.left,c)
			t.push(p.right,c)
			break
		}
		if p.left != nil {
			t.push(p.left,c)
			break
		}
		if p.right != nil {
			t.push(p.right,c)
			break
		}
	}
}
func(t *Tree) push(node *Node, c chan *Node) {
	c <- node
}
func (t *Tree) FindMin() (key int) {
	return findMin(t.root)
}
func findMin(node *Node) int {
	if node.left == nil {
		return node.key
	} else {
		node = node.left
		return findMin(node)
	}
}

func (t *Tree) FindMax() (key int) {
	return findMax(t.root)

}
func findMax(node *Node) int {
	if node.right == nil {
		return node.key
	} else {
		node = node.right
		return findMax(node)
	}
}
func (t *Tree) DeleteMin() {
	t.deleteMin(t.root)
}
func (t *Tree) deleteMin(node *Node) {
	if node.left == nil {
		node = nil
		return
	} else {
		//fmt.Println("1")
		node = node.left
		t.deleteMin(node)
	}
}
func (t *Tree) DeleteMax() {

}
func (t *Tree) Delete(key int) bool{
	if t.Contain(key) == false {
		fmt.Println("the key did not exist")
		return false
	}
	return t.delete(t.root, key)
}
func (t *Tree) delete(node *Node,target int) bool {
	if node.left == nil && node.right == nil  {
		return true
	}
	fmt.Println("count")

	//已发现目标，node的左或右子节点是目标
	if node.left.key == target || node.right.key == target{
		//node左子节点是目标
		if node.left.key == target {
			fmt.Println("leftttt")
			//左子节点有左右子树
			if node.left.left != nil && node.left.right != nil {
				tail := node.left.left
				node.left = node.left.right
				cur := node.left
				for i := 0; i < t.count;i ++ {

					if cur.left != nil {
						cur = cur.left
					} else {
						cur.left = tail

						break
					}

				}
			}
			//左子节点有左子树无右子树
			if node.left.right == nil && node.left.left != nil{
				node.left = node.left.left
			}
			//左子节点有右子树而无左子树
			if node.left.right == nil && node.left.left != nil {
				node.left = node.left.right

			}
			//左子节点无左子树和右子树
			if node.left.right == nil && node.left.left == nil {
				node.left = nil
			}
		}

		//node右子节点是目标
		if node.right.key == target {
			fmt.Println("rightttt")
			//右子节点有左右子树
			if node.right.left != nil && node.right.right != nil {
				fmt.Println("r1")
				tail := node.right.left
				node.right = node.right.right
				cur := node.right
				for i := 0; i < t.count; i ++ {
					if cur.left != nil {
						cur = cur.left
					} else {
						cur.left = tail
						break
					}
				}
				return true
			}
			//右子节点有右子树没有左子树
			if node.right.left == nil && node.right.right != nil {
				fmt.Println("r2")
				//fmt.Println(node.right.right)
				fmt.Println("nd",node)
				node.right = node.right.right
				fmt.Println("nooood",node)

				fmt.Println("roooot",t.root)
				fmt.Println("rooot.left",t.root.left)
				return true
			}
			//右子节点有左子树没有右子树
			if node.right.left != nil && node.right.right == nil {
				fmt.Println("r3")
				node.right = node.right.left
				return true
			}
			//右子节点没有左右子树
			if node.right.left == nil && node.right.right == nil {
				fmt.Println("r4")
				node.right = nil
				return true
			}

		}
	} else {
		node = node.left
		t.delete(node, target)
		node = node.right
		fmt.Println("node...", node)
		t.delete(node, target)
	}

	return true

}
func main() {
	var t Tree
	//t.Initial()
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


/*
	for m := 0; m < 55; m ++ {
		a := t.Contain(m)
		if a == true {
			fmt.Println(m)

		}
	}
	*/
	/*
	fmt.Println(t.Contain(6))  // true
	fmt.Println(t.Contain(29)) // false
	fmt.Println(t.Search(6))   // 60
	fmt.Println("root",t.root)
	t.PreOrder()
	fmt.Println(t.count)
	t.LevelOfChan()
	min := t.FindMin()
	fmt.Println("min:",min)//min:0
	max := t.FindMax()
	fmt.Println("max:",max)// max:47
	*/
	fmt.Println("start delete")
	t.Delete(18)
	fmt.Println("after delete 18")
	t.PreOrder()
	//fmt.Println("222")
	//t.LevelOfChan()
	/*
	t.DeleteMin()
	fmt.Println("d min:", t.FindMin())
	*/

	fmt.Println("final root",t.root)
}
