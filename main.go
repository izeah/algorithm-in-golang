package main

import (
	"algorithm-in-golang/recursion"
	"algorithm-in-golang/searching"
	"algorithm-in-golang/sorting"
	"fmt"
)

func main() {
	// recursive algorithm
	fmt.Println("RECURSIVE FUNCTION")
	fmt.Println("===================")
	recursion.Run("C:/Users/Public")

	// given sorted array of integer
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	fmt.Println("\n\nLINEAR SEARCHING")
	fmt.Println("=========================")
	linear_res := searching.LinearSearch(arr, 7)
	fmt.Println(linear_res)

	fmt.Println("\n\nBINARY SEARCHING")
	fmt.Println("=========================")
	binary_res := searching.BinarySearch(arr, 7)
	fmt.Println(binary_res)

	// given unsorted array of integer
	randomArrInt := []int{4235, 10203, 283, 10, 3893, 28, 48, 1, 493, 3, 489, 97, 1003, 2849}

	fmt.Println("\n\nBUBBLE SORT")
	fmt.Println("=========================")
	fmt.Println("Before")
	fmt.Println(randomArrInt)

	// function bubble sort
	sorting.BubbleSort(randomArrInt)

	fmt.Println("After")
	fmt.Println(randomArrInt)

	// given unsorted array of integer
	randomArrInt = []int{4235, 10203, 283, 10, 3893, 28, 48, 1, 493, 3, 489, 97, 1003, 2849, 2342}

	fmt.Println("\n\nMERGE SORT")
	fmt.Println("=========================")
	fmt.Println("Before")
	fmt.Println(randomArrInt)

	// function merge sort
	sortedArrInt := sorting.MergeSort(randomArrInt)

	fmt.Println("After")
	fmt.Println(sortedArrInt)

	// given unsorted array of integer
	randomArrInt = []int{4235, 10203, 283, 10, 3893, 28, 48, 1, 493, 3, 489, 97, 1003, 2849, 2342}

	fmt.Println("\n\nSELECTION SORT")
	fmt.Println("=========================")
	fmt.Println("Before")
	fmt.Println(randomArrInt)

	// function selection sort
	sorting.SelectionSort(randomArrInt)

	fmt.Println("After")
	fmt.Println(randomArrInt)

	// given unsorted array of integer
	randomArrInt = []int{4235, 10203, 283, 10, 3893, 28, 48, 1, 493, 3, 489, 97, 1003, 2849, 2342}

	fmt.Println("\n\nINSERTION SORT")
	fmt.Println("=========================")
	fmt.Println("Before")
	fmt.Println(randomArrInt)

	// function selection sort
	sorting.InsertionSort(randomArrInt)

	fmt.Println("After")
	fmt.Println(randomArrInt)
}
