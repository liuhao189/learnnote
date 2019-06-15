package main

import "fmt"

func printMsg(msg string) {
	fmt.Printf("%s\n", msg)
}

func printSlice(x []int) {
	fmt.Printf("len = %d, cap = %d, slice=%v\n", len(x), cap(x), x)
}

func main() {
	s := []int{1, 2, 3}
	arr := []int{4, 5, 6}
	s2 := arr[:]
	s3 := arr[2:3]
	printSlice(s)
	printSlice(arr)
	printSlice(s2)
	printSlice(s3)
	var numbers []int
	printSlice(numbers)
	if numbers == nil {
		printMsg("Slice is empty!")
	}

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	fmt.Println("numbers == ", numbers)

	fmt.Println("numbers[1:4] == ", numbers[1:4]) //1,2,3

	fmt.Println("numbers[:3] == ", numbers[:3]) //0,1,2

	fmt.Println("numbers[4:] == ", numbers[4:]) //4,5,6,7,8

	numbers1 := make([]int, 0, 5)

	printSlice(numbers1)

	number2 := numbers[:2] //0,1

	printSlice(number2)

	numbers3 := numbers[2:5] //2,3,4

	printSlice(numbers3)

	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	copy(numbers1, numbers)

	printSlice(numbers1)
}
