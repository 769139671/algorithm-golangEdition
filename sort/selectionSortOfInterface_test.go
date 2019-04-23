package main

import (
	"math/rand"
	"testing"
)
func GenerateArr(n int,t string) interface{}{
	switch t {
	case "int":
		b := []int{}
		for i := 0; i < n; i ++ {
			a := rand.Int()
			//fmt.Println(a)
			b = append(b,a)
		}
		return  b

	case "float64":
		b := []float64{}
		for i := 0; i < n; i ++ {
			a := rand.Float64()
			//fmt.Println(a)
			b = append(b,a)
		}
		return  b

	case "string":
		b := []string{}
		for i := 0; i < n; i ++ {
			a := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
			//fmt.Println(a)
			//v := []string{}
			for _, x := range a {
				b = append(b, x)

			}

		}
		return  b
	}

	return nil
}

func TestSort(t *testing.T) {
	n := 100000
	//z := GenerateArr(n)
	//b := z.([]float64)
	tp := "int"
	z2 := GenerateArr(n,tp)
	b := z2.([]int)
	//fmt.Println(b)
	c := Sort(b)
	//fmt.Println(c)
	switch c.(type) {
	case []float64:
		e := c.([]float64)
		for i := 0; i <len(e); i ++ {
			for j := i +1; j <len(e); j ++ {
				if e[i] > e[j] {
					t.Errorf("wrone")
				}
			}
		}
		case []int:
		e := c.([]int)
		for i := 0; i <len(e); i ++ {
			for j := i +1; j <len(e); j ++ {
				if e[i] > e[j] {
					t.Errorf("wrone")
				}
			}
		}
	}
}