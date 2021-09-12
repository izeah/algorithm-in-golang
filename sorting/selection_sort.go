package sorting

import "fmt"

func SelectionSort(data []int) {
	fmt.Println("Process")
	length := len(data)
	for i := 0; i < length; i++ {
		maxIndex := 0

		// proses nyari index dengan nilai element tertinggi
		for j := 1; j < length-i; j++ {
			if data[j] > data[maxIndex] {
				maxIndex = j
			}
		}

		// menampung nilai elemen tertinggi ke index array terakhir
		data[length-i-1], data[maxIndex] = data[maxIndex], data[length-i-1]
		fmt.Println(data)
	}
}
