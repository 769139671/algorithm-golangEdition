package main

import "fmt"

type IndexHeap struct {
	i []int
	heap []int
	reverse []int
}
func (hi *IndexHeap) Initial() {
	hi.heap = append(hi.heap, 0)
	hi.i = append(hi.i,0)
}
func (hi *IndexHeap) Insert(v int) bool {
	//fmt.Println(len(hi.heap))
	hi.i = append(hi.i,len(hi.heap))
	hi.heap = append(hi.heap, v)
	j := len(hi.heap) - 1
	hi.indexShiftUp(hi.heap,j)
	return true
}
func (hi *IndexHeap) Delete(i int) bool {
	//索引i范围为：[1,len(h.heap))
	if i < 1 || i > len(hi.heap) -1{
		return false
	}
	if len(hi.heap) <=2 {
		hi.heap = hi.heap[:len(hi.heap) - 1]
		hi.i = hi.i[:len(hi.i)-1]
		return true
	} else {
		//fmt.Println("before change:",hi.heap)
		//fmt.Println("before change i:",hi.i)
		//hi.i[len(hi.i)-1],hi.i[i] = hi.i[i], hi.i[len(hi.i)-1]
		//fmt.Println("change:",hi.heap)
		//fmt.Println("change i:",hi.i)

		//hi.heap = hi.heap[:len(hi.heap)-1]
		fmt.Println("before delete 1",hi.heap)
		fmt.Println("hi.i:",hi.i)
		fmt.Println("hi.i[i] :", hi.i[i])
		hi.heap = append(hi.heap[:hi.i[i]],hi.heap[hi.i[i]+1:]...)
		fmt.Println("delete 1:",hi.heap)
		hi.i = hi.i[:len(hi.i)-1]
		fmt.Println("new i:",hi.i)
		fmt.Println("new heap:",hi.heap)
		for pi,p := range hi.i {
			if p == i {
				hi.i = append(hi.i[:pi],hi.i[pi+1:]...)

			}
		}
		if len(hi.heap) <= 2 {
			return true
		} else {
			hi.indexShiftDown(hi.heap,i, len(hi.heap))
		}
	}
	return true
}
func (hi *IndexHeap) indexShiftDown(arr []int,i int, l int) {
	if l <= 3 {
		if arr[i] < arr[2*i] {
			hi.i[i], hi.i[2*i] = hi.i[2*i], hi.i[i]
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
				hi.i[i], hi.i[2*i] = hi.i[2*i], hi.i[i]
				z := 2*i
				//fmt.Println("a1")
				hi.indexShiftDown(arr, z,l)
			} else {
				hi.i[i], hi.i[2*i +1] = hi.i[2*i +1], hi.i[i]
				z := 2*i+1
				hi.indexShiftDown(arr, z,l)
			}
		}
	} else {
		//只有左子树
		if i * 2 <= l -1  {

			if arr[i] >= arr[i*2] {
				return
			} else {
				hi.i[i], hi.i[i*2] = hi.i[i*2], hi.i[i]
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
func (hi *IndexHeap)indexShiftUp(arr []int, j int)  {
	if j <= 1 {
		return
	} else {
		if arr[j] > arr[j/2] {
			hi.i[j],hi.i[j/2] = hi.i[j/2], hi.i[j]
			j = j/2
			hi.indexShiftUp(arr,j)
		} else {
			return
		}
	}
	return
}
func main() {
	var a IndexHeap
	a.Initial()
	arr := []int{3,7,5,1,9,2,8,5,6,1,15,4}
	for  _,x := range arr {
		a.Insert(x)
	}
	fmt.Println("heap:",a)
	a.Delete(1)
	fmt.Println("delete:",a)


}
