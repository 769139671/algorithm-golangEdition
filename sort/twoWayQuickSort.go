package main

import (
	"math/rand"
)

func TwoWayQuickSort(arr []int) {
	if len(arr) <=1 {
		return
	}
	x := rand.Intn(len(arr))
	arr[0], arr[x] = arr[x], arr[0]
	l := 0
	//r := len(arr) - 1
	m := 1
	n := len(arr) - 1
	for i := 0; i < len(arr); i ++ {
		if m > n {
			break
		}
		//把=arr[l]的元素均匀的分配到前后两个数组，实现归并树的平衡
		if arr[m] >= arr [l] && arr[n] <= arr[l] {
			arr[m], arr[n] = arr[n], arr[m]
			m ++
			n --
		} else {
				if arr[m] > arr[l] {
					n--
				} else {
						m ++
				}

		}
	}
	//fmt.Println(n)
	arr[l], arr[n] = arr[n], arr[l]
	//fmt.Println(arr[:n])
	//fmt.Println(arr[n + 1:])
	TwoWayQuickSort(arr[:n])
	TwoWayQuickSort(arr[n + 1:])

}
/*
func main() {
	a := []int{5,4,8,10,5,4,1,3,5,9,2,1,1,8,8,10,15,21,15}
	TwoWayQuickSort(a)
	fmt.Println(a)
}
*/
