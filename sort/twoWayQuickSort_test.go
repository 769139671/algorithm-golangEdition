package main

import (
	"math/rand"
	"testing"
)

func TestTwoWayQuickSort(t *testing.T) {

	arr := []int{}
	for i := 0; i < 100000; i ++ {
		b := rand.Intn(10)
		arr = append(arr, b)
	}
	//fmt.Println(arr)

	TwoWayQuickSort(arr)
	for i := 0; i <len(arr); i ++ {
		for j := i +1; j <len(arr); j ++ {
			if arr[i] > arr[j] {
				t.Errorf("wrone")
			}
		}
	}




/*
	QuickSort(arr)
	for i := 0; i <len(arr); i ++ {
		for j := i +1; j <len(arr); j ++ {
			if arr[i] > arr[j] {
				t.Errorf("wrone")
			}
		}
	}
*/

}
