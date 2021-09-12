package sorting

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) / 2

	// divide and conquer left side sub-array recursively
	left := MergeSort(data[:middle])
	// divide and conquer right side sub-array recursively
	right := MergeSort(data[middle:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	// initiate new sorted array
	result := make([]int, len(left)+len(right))
	for i := 0; len(left) > 0 || len(right) > 0; i++ {
		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				// input element dari sub array ke array result
				result[i] = left[0]
				// hapus element yang sudah diinput ke sorted array result
				left = left[1:]
			} else {
				// input element dari sub array ke array result
				result[i] = right[0]
				// hapus element yang sudah diinput ke sorted array result
				right = right[1:]
			}
		} else if len(left) > 0 {
			// input element dari sub array ke array result
			result[i] = left[0]
			// hapus element yang sudah diinput ke sorted array result
			left = left[1:]
		} else if len(right) > 0 {
			// input element dari sub array ke array result
			result[i] = right[0]
			// hapus element yang sudah diinput ke sorted array result
			right = right[1:]
		}
	}
	return result
}
