package main

func InsertionSort(arr []int) []int {
	arrInt := arr
	for i := 1; i < len(arrInt); i ++ {
		cur := i
		for j := cur - 1; j >= 0 ; j -- {

			if arrInt[cur] < arrInt[j] {
				arrInt[j], arrInt[cur] = arrInt[cur], arrInt[j]
				cur = j
			} else {
				break
			}

		}
	}
	return arrInt


}
/*
func main() {
	arr := []int{13,5,4,6,9,8,1,7,2}

	fmt.Println(bubbleSort(arr))
}
*/

