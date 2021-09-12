package sorting

import "fmt"

func BubbleSort(data []int) {
	fmt.Println("Process")
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			// current element < previous element
			if data[j] < data[j-1] {
				// swap element // tukar current element dan previous element
				data[j], data[j-1] = data[j-1], data[j]
				fmt.Println(data)
			}
		}
	}
}
