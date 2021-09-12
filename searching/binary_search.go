package searching

func BinarySearch(arr []int, x int) int {
	var (
		start = 0
		end   = len(arr) - 1
		mid   = -1
	)

	for start != mid && end != mid {
		mid = start + (end-start)/2
		if arr[mid] == x {
			return mid
		}

		if arr[mid] > x {
			end -= 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
