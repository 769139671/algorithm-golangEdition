package main

func find(arr []int) int {
	n := 0
	for i := 0; i <len(arr)-1; i ++ {
		for j := i +1; j <len(arr); j ++ {
			if arr[i] > arr[j] {
				n ++
			}
		}
	}
	return n
}
/*
func main() {
	a := []int{3,7,1,5,6,2,8,4}
	b := []int{1,2,3,4,5,6,7,8}
	fmt.Println(find(a))
	fmt.Println(find(b))
}
*/
