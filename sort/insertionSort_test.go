package main

import (
	"math/rand"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	n := 100000
	b := []int{}
	for i := 0; i < n; i ++ {
		a := rand.Int()
		//fmt.Println(a)
		b = append(b,a)
	}
	c := InsertionSort(b)
	//fmt.Println(c)

	for i := 0; i <len(c); i ++ {
		for j := i +1; j <len(c); j ++ {
			if c[i] > c[j] {
				t.Errorf("wrone")
			}
		}
	}
}