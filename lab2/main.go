package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int = 10

	fmt.Print("Number of integers to sort: ")
	fmt.Scanf("%d\n", &n)

	rand.Seed(time.Now().Unix())
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(1000)
	}

	var start time.Time
	var duration time.Duration
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = arr[i]
	}

	fmt.Println("Running Quick Sort ...")
	start = time.Now()
	QuickSort(a)
	duration = time.Since(start)
	fmt.Printf("It took %d ns to sort the array of %d integers.\n\n", duration.Nanoseconds(), n)

	for i := 0; i < n; i++ {
		a[i] = arr[i]
	}

	fmt.Println("Running Merge Sort ...")
	start = time.Now()
	MergeSort(a)
	duration = time.Since(start)
	fmt.Printf("It took %d ns to sort the array of %d integers.\n\n", duration.Nanoseconds(), n)

	for i := 0; i < n; i++ {
		a[i] = arr[i]
	}

	fmt.Println("Running Heap Sort ...")
	start = time.Now()
	HeapSort(a)
	duration = time.Since(start)
	fmt.Printf("It took %d ns to sort the array of %d integers.\n\n", duration.Nanoseconds(), n)

	for i := 0; i < n; i++ {
		a[i] = arr[i]
	}

	fmt.Println("Running Bubble Sort ...")
	start = time.Now()
	BubbleSort(a)
	duration = time.Since(start)
	fmt.Printf("It took %d ns to sort the array of %d integers.\n\n", duration.Nanoseconds(), n)

	for i := 0; i < n; i++ {
		a[i] = arr[i]
	}

	fmt.Println("Running Selection Sort ...")
	start = time.Now()
	SelectionSort(a)
	duration = time.Since(start)
	fmt.Printf("It took %d ns to sort the array of %d integers.\n\n", duration.Nanoseconds(), n)

	for i := 0; i < n; i++ {
		a[i] = arr[i]
	}

	fmt.Println("Running Insertion Sort ...")
	start = time.Now()
	InsertionSort(a)
	duration = time.Since(start)
	fmt.Printf("It took %d ns to sort the array of %d integers.\n\n", duration.Nanoseconds(), n)
}

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) / 2
	left := MergeSort(data[:middle])
	right := MergeSort(data[middle:])

	return merge(left, right)
}
func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	for i := 0; len(left) > 0 || len(right) > 0; i++ {
		if len(left) > 0 && len(right) > 0 {
			if left[0] < right[0] {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
		} else if len(left) > 0 {
			result[i] = left[0]
			left = left[1:]
		} else if len(right) > 0 {
			result[i] = right[0]
			right = right[1:]
		}
	}
	return result
}

func QuickSort(nums []int) {
	recursionSort(nums, 0, len(nums)-1)
}
func recursionSort(data []int, left int, right int) {
	if left < right {
		pivot := partition(data, left, right)
		recursionSort(data, left, pivot-1)
		recursionSort(data, pivot+1, right)
	}
}
func partition(data []int, left int, right int) int {
	for left < right {
		for left < right && data[left] <= data[right] {
			right--
		}
		if left < right {
			data[left], data[right] = data[right], data[left]
			left++
		}

		for left < right && data[left] <= data[right] {
			left++
		}
		if left < right {
			data[left], data[right] = data[right], data[left]
			right--
		}
	}
	return left
}

func HeapSort(data []int) []int {
	heapify(data)
	for i := len(data) - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]
		siftDown(data, 0, i)
	}
	return data
}
func heapify(data []int) {
	for i := (len(data) - 1) / 2; i >= 0; i-- {
		siftDown(data, i, len(data))
	}
}
func siftDown(heap []int, lo, hi int) {
	root := lo
	for {
		child := root*2 + 1
		if child >= hi {
			break
		}
		if child+1 < hi && heap[child] < heap[child+1] {
			child++
		}
		if heap[root] < heap[child] {
			heap[root], heap[child] = heap[child], heap[root]
			root = child
		} else {
			break
		}

	}
}

func BubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}

func SelectionSort(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		maxIndex := 0
		for j := 1; j < length-i; j++ {
			if data[j] > data[maxIndex] {
				maxIndex = j
			}
		}
		data[length-i-1], data[maxIndex] = data[maxIndex], data[length-i-1]
	}
}

func InsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			j := i - 1
			temp := data[i]
			for j >= 0 && data[j] > temp {
				data[j+1] = data[j]
				j--
			}
			data[j+1] = temp
		}
	}
}
