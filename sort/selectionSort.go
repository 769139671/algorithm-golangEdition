package main

func selectionSort(arr []int) (b []int){

	for i := 0; i < len(arr); i ++ {
		minIndex := i
		for j := i+1; j < len(arr); j ++ {

			if arr[minIndex] > arr[j] {
				arr[minIndex],arr[j] = arr[j],arr[minIndex]
			}
		}
	}
	return arr
}
/*
func main() {
	arr := []int{10,9,8,7,6,5,4,3,2,1}
	fmt.Println(arr)
	fmt.Println(selectionSort(arr))

}
*/
