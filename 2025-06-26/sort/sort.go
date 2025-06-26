package sort

// BubbleSort sorts a slice of integers in ascending order using the bubble sort algorithm.
// It modifies the input slice in place.
func BubbleSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break // optimization: stop if no swaps in a pass
		}
	}
}
