package main

import "fmt"

func main() {
	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	var arr = [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	for i = 0; i < 3; i++ {
		for j = 0; j < 4; j++ {
			fmt.Printf("arr[%d][%d] = %d \n", i, j, arr[i][j])
		}
	}

	var balance = []int{1000, 2, 3, 17, 50}
	var avg float32
	avg = getAverage(balance, 5)
	fmt.Printf("Average is %f.\n", avg)
}

func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	avg = float32(sum / size)
	return avg
}
