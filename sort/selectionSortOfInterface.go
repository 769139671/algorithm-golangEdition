package main

import "fmt"

func Sort(arr interface{}) interface{} {
	switch  arr.(type) {
	case []int:
		//fmt.Println("is []int")
		b := arr.([]int)
		for i := 0; i < len(b); i ++ {
			minIndex := i
			for j := i +1; j < len(b); j ++ {
				if b[minIndex] > b[j] {
					b[minIndex], b[j] = b[j], b[minIndex]
				}
			}
		}
		return b

	case []float64:
		fmt.Println("is []float64")
		b := arr.([]float64)
		for i := 0; i <len(b); i ++ {
			minIndex := i
			for j := i+1; j < len(b); j ++ {
				if b[minIndex] > b[j] {
					b[minIndex], b[j] = b[j], b[minIndex]
				}
			}
		}
		return b

	case []string:
		fmt.Println("is []string")
		b := arr.([]string)
		for i := 0; i <len(b); i ++ {
			minIndex := i
			for j := i+1; j < len(b); j ++ {
				if b[minIndex] > b[j] {
					b[minIndex], b[j] = b[j], b[minIndex]
				}
			}
		}
		return b
	}
	return nil
}
/*
func main() {

	a1 := []int{10,9,8,7,6,5,4,3,2,1}
	fmt.Println(Sort(a1))

	a2 := []string{"c","d","b","a","e","z"}

	fmt.Println(Sort(a2))
	a3 := []float64{5.22, 6.33, 1.44, 0.25, 0.79, 8.77, 9.65}
	fmt.Println(Sort(a3))
}
*/

