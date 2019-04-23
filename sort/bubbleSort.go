package main

func bubbleSort(arr []int)  []int {

	for i := 0; i < len(arr) - 1; i ++ {
		for j := 0; j < len(arr) -1 ; j ++ {
			if arr[j] > arr[j +1] {
				arr[j], arr[j +1] = arr[j + 1], arr[j]
			}
		}

	}
	return arr

}
/*
func main() {
	arr := []int{13,5,4,6,9,8,1,7,2}

	fmt.Println(bubbleSort(arr))
}
*/