package main

import (
	"math/rand"
	"testing"
)
/*
func TestThreeWayQuickSort(t *testing.T) {
	arr := []int{}
	for i := 0; i < 100000; i ++ {
		b := rand.Int()
		arr = append(arr, b)
	}
	//fmt.Println(arr)



	ThreeWayQuickSort(arr)

	for i := 0; i <len(arr); i ++ {
		for j := i +1; j <len(arr); j ++ {
			if arr[i] > arr[j] {
				t.Errorf("wrone")
			}
		}
	}
}
*/
func BenchmarkThreeWayQuickSort(b *testing.B) {
	//b.StopTimer()
	arr := []int{}
	for i := 0; i < 10000000; i ++ {
		b := rand.Int()
		arr = append(arr, b)
	}
	//b.StartTimer()
	ThreeWayQuickSort(arr)
	//b.StopTimer()


}
