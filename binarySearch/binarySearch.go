package main

import "fmt"

func find(target int, arr []int) int{
	l := 0
	r := len(arr) - 1

	for i := 0; i <len(arr); {
		if l == r-1 && r == l +1{
			if arr[r] == target {
				return r
			} else {
				if arr[l] == target{
					return l
				}
				fmt.Println("not find")
				return 0
			}
		}
		mid := (l + r)/2
		if arr[mid] == target {
			return mid
		} else {
			if target> arr[mid] {
				l = mid
			} else {
				r = mid
			}
		}
	}
	fmt.Println("not find")
	return 0

}
/*
func main() {

	arr := []int{1,2,3,3,3,3,4,5,6}
	fmt.Println(find(1,arr))


}
*/