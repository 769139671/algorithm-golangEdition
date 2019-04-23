package main

import "fmt"

func ThreeWayQuickSort(arr []int) {
	if len(arr) <=  1{
		return
	}
	l := 0
	m := 1
	//<arr[l]队列的最后一个元素索引+1，是=arr[l]的第一个元素
	lt := 1
	//>arr[l]队列的第一个元素索引-1，是未知元素
	gt := len(arr)-1
	for i := 0; i < len(arr); i ++ {
		if m > gt {
			break
		}
		if arr[m] < arr[l] {
			arr[m], arr[lt] = arr[lt], arr[m]
			lt ++
			m ++
		} else {
			if arr[m] == arr[l] {
				//arr[lt+1]是=arr[l]队列的第一个元素
				//arr[lt +1], arr[i] = arr[i],arr[lt + 1]
				m ++
			} else {
				//gt是未知元素，i 不自增
				arr[m], arr[gt] = arr[gt], arr[m]
				gt --

			}
		}
	}
	//fmt.Println(arr[l])
	//fmt.Println(arr[lt-1])
	arr[l], arr[lt-1] = arr[lt-1], arr[l]
	//fmt.Println(arr[gt + 1])
	/*
	for i < len(arr){
		if i > gt {
			break
		}
		if arr[i] < arr[l] {

			arr[i], arr[lt] = arr[lt], arr[i]
			lt ++
			i ++
		} else {
				if arr[i] == arr[l] {
					//arr[lt+1]是=arr[l]队列的第一个元素
					//arr[lt +1], arr[i] = arr[i],arr[lt + 1]
					i ++
				} else {
						//gt是未知元素，i 不自增
						arr[i], arr[gt] = arr[gt], arr[i]
						gt --

				}
		}

	}
	*/
	ThreeWayQuickSort(arr[:lt-1])
	ThreeWayQuickSort(arr[gt + 1:])


}
func main() {
	a := []int{5,4,8,10,5,4,1,3,5,9,2,1,1,8,8,10,15,21,15}
	ThreeWayQuickSort(a)
	fmt.Println(a)
}
