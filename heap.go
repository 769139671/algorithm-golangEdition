package main

import (
	"fmt"
	"math/rand"
)

type Heap struct {
	heap []int
}
//初始化堆，将索引为0的位置元素定义为0
func (h *Heap) Initial() {
	h.heap = append(h.heap, 0)
}
//实现堆中插入新元素元素
func (h *Heap) Insert(v int) bool {
	h.heap = append(h.heap, v)
	j := len(h.heap) - 1
	shiftUp(h.heap,j)
	return true
}
//实现删除索引为i的元素
func (h *Heap) Delete(i int) bool{
	//索引i范围为：[1,len(h.heap))
	if i < 1 || i > len(h.heap) -1{
		return false
	}
	if len(h.heap) <=2 {
		h.heap = h.heap[:len(h.heap) - 1]
		return true
	} else {
		h.heap[len(h.heap)-1],h.heap[i] = h.heap[i], h.heap[len(h.heap)-1]
		h.heap = h.heap[:len(h.heap)-1]
		if len(h.heap) <= 2 {
			return true
		} else {
			shiftDown(h.heap,i, len(h.heap))
		}
	}
	return true
}
//提取最大值，返回值为int
func (h *Heap) ExtractMax() int {
	type IndexMax struct {
		i int
		heap []int
	}
	result := h.heap[1]
	h.Delete(1)
	return result
}

//优化的堆
func (h *Heap) Heapify()  {
	l := (len(h.heap)-1) / 2
	for i := l; i >0; i -- {
		shiftDown(h.heap,i,len(h.heap))
	}

}
//实现堆排序
func (h *Heap) HeapSort(arr []int) {

	for _,a := range arr {
		h.heap = append(h.heap, a)
	}
	l := len(h.heap)
	h.Heapify()

	for i := l -1; i >0; {
		h.heap[1], h.heap[i] = h.heap[i], h.heap[1]
		i --
		var e Heap
		e.heap = h.heap[:i+1]
		e.Heapify()
		//fmt.Println("sort",h.heap)
	}
}

//元素下移
func shiftDown(arr []int, i int, l int) {
	if l <= 3 {
		if arr[i] < arr[2*i] {
			arr[i], arr[2*i] = arr[2*i], arr[i]
			return
		} else {
			return
		}
	}
	//有左右子树
	if i * 2 +1 <= l -1{

		if arr[i] >= arr[2*i] && arr[i] >= arr[2*i +1] {

			return
		}else {
			if arr[2*i] >= arr[2*i+1] {
				arr[i], arr[2*i] = arr[2*i], arr[i]
				z := 2*i
				//fmt.Println("a1")
				shiftDown(arr, z,l)
			} else {
				arr[i], arr[2*i +1] = arr[2*i +1], arr[i]
				z := 2*i+1
				shiftDown(arr, z,l)
			}
		}
	} else {
		//只有左子树
		if i * 2 <= l -1  {

			if arr[i] >= arr[i*2] {
				return
			} else {
				arr[i], arr[i*2] = arr[i*2], arr[i]
				return
			}
		} else {
			//没有子树
			if i*2 >l -1 && i* 2+1 >l-1 {
				return
			}
		}
	}
}
//元素上移
func shiftUp(arr []int, j int)  {
	if j <= 1 {
		return
	} else {
		if arr[j] > arr[j/2] {
			arr[j],arr[j/2] = arr[j/2], arr[j]
			j = j/2
			shiftUp(arr,j)
		} else {
			return
		}
	}
	return
}
/*
func main() {
	var a Heap
	a.Initial()
	fmt.Println(a.heap)
	for i := 0; i <20; i ++ {
		a.Insert(rand.Intn(20))
	}
	fmt.Println("heap",a.heap)

	for j :=0; j <20; j ++ {
		//a.Delete(1)
		fmt.Println("max",a.ExtractMax())
	}

	var b Heap
	b.Initial()
	fmt.Println(b.heap)
	arr := []int{3,7,5,1,9,2,8,5,6,1,15,4}
	for _,x := range arr {
		b.heap = append(b.heap,x)
	}
	b.Heapify()
	fmt.Println("b heap of heapify",b.heap)

	var d Heap
	arr = []int{3,7,5,1,9,2,8,5,6,1,15,4}
	d.Initial()
	d.HeapSort(arr)
	fmt.Println("sort",d.heap)

	var e Heap
	e.Initial()

}
*/