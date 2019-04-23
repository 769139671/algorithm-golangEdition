package main

import (
	"math/rand"
)

func RandomQuickSort(arr []int) {


	//
	//
	//fmt.Println(x)
	//fmt.Println(len(arr))

	if len(arr) <= 1 {
		return
	}

	r := len(arr) - 1
	j := 1
	x := rand.Intn(len(arr))
	arr[0], arr[x] = arr[x], arr[0]
	l := 0
	key := arr[l]
	for i := 1; i <= r; {
		if arr[i] < key {
			arr[i], arr[j] = arr[j], arr[i]
			i ++
			j ++
		} else {
			i ++
		}
	}
	arr[l],arr[j-1] = arr[j-1], arr[l]
	RandomQuickSort(arr[:j-1])
	RandomQuickSort(arr[j:])

}
/*
func main() {
	a := []int{4,8,10,1,3,9,2,5}
	RandomQuickSort(a)
	//fmt.Println(a)
}
*/
