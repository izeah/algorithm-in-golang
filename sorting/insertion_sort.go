package sorting

import "fmt"

func InsertionSort(data []int) {
	fmt.Println("Process")
	for i := 1; i < len(data); i++ {
		// current element < previous element
		if data[i] < data[i-1] {
			j := i - 1
			// menampung nilai yang mau dipindahin
			temp := data[i]

			// geser element
			for j >= 0 && data[j] > temp {
				data[j+1] = data[j]
				j--
			}

			// update nilai yang ditampung barusan
			data[j+1] = temp
			fmt.Println(data)
		}
	}
}
