package main

func QuickSort(arr []int) {

	l := 0
	r := len(arr) - 1
	j := 1
	if len(arr) <= 1 {
		return
	}
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
	QuickSort(arr[:j-1])
	QuickSort(arr[j:])

}
/*
func main() {
	a := []int{4,8,10,1,3,9,2,5}
	QuickSort(a)
	fmt.Println(a)
}
*/