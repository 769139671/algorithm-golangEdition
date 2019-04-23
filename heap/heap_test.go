package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestHeap_ExtractMax(t *testing.T) {
	var te Heap
	te.Initial()
	for i := 0; i < 100000; i ++ {
		j := rand.Int()
		te.Insert(j)

	}
	fmt.Println("heap:",te.heap)
	/*
	for j := 0; j<10; j ++ {
		fmt.Println(te.ExtractMax())
	}
	*/

	sort := []int{}
	l := len(te.heap)
	fmt.Println("len",l)
	//b := te.ExtractMax()
	//fmt.Println(b)
	//fmt.Println("ex",te.ExtractMax())
	//testing.B.StartTimer(te.ExtractMax())

	for k := 0; k<l-1 ; k ++ {
		sort = append(sort, te.ExtractMax())
	}
	//fmt.Println("sort",sort)

	for m := 1; m <len(sort);  m++{
		for n := m +1; n <len(sort); n++ {
			if sort[m] < sort[n] {
				t.Errorf("wrone")
			}

		}

	}



}
func BenchmarkHeap_HeapSort(b *testing.B) {
	b.StopTimer()
	var arr []int

	for i := 0; i < 100000; i ++ {
		j := rand.Int()
		arr = append(arr,j)

	}
	//b.StartTimer()
	var te Heap
	te.Initial(&te)
	b.StartTimer()
	te.HeapSort(arr)
	//fmt.Println(te.heap)

}