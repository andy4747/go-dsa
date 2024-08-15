package arrays

func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		space := i
		temp := arr[i]
		for space > 0 && arr[space-1] > temp {
			arr[space] = arr[space-1]
			space = space - 1
		}
		arr[space] = temp
	}
}
