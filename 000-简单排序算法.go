package main

import "fmt"

func main() {
	arr := []int{9, 7, 3, 5, 1, 15, 11}

	//BubbleSort(arr)
	//SelectSort(arr)
	QuickSort(arr, 0, 6)
	//InsertSort(arr)
	//ShellSort(arr)
	//MergeSort(arr)
	fmt.Println(arr)
}

func BubbleSort(arr []int) {
	var isChange bool
	for j := 0; j < len(arr)-1; j++ {
		isChange = false
		for i := 0; i < len(arr)-j-1; i++ {
			if arr[i] < arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isChange = true
			}
		}
		if !isChange {
			break
		}
	}
}

func SelectSort(arr []int) {
	var max int
	for i := 0; i < len(arr)-1; i++ {
		max = i
		for j := i + 1; j < len(arr); j++ {
			if arr[max] < arr[j] {
				max = j
			}
		}
		arr[i], arr[max] = arr[max], arr[i]
	}
}

func QuickSort(arr []int, l, r int) {
	i, j := l, r
	pivot := arr[(l+r)/2]
	for i <= j {
		for pivot > arr[i] {
			i++
		}
		for pivot < arr[j] {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	if l < j {
		QuickSort(arr, l, j)
	}
	if i < r {
		QuickSort(arr, i, r)
	}
}

func InsertSort(arr []int) {
	var temp int
	for i := 1; i < len(arr); i++ {
		temp = arr[i]
		j := i - 1
		for j >= 0 && temp < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}

func ShellSort(arr []int) {
	for step := len(arr) / 2; step > 0; step /= 2 {
		for i := step; i < len(arr); i++ {
			j := i
			temp := arr[j]
			for j-step >= 0 && arr[j-step] > temp {
				arr[j] = arr[j-step]
				j = j - step
			}
			arr[j] = temp
		}
	}
}

func MergeSort(arr []int) {
	temp := make([]int, len(arr))
	sort(arr, temp, 0, len(arr)-1)
}

func sort(arr, temp []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		sort(arr, temp, left, mid)
		sort(arr, temp, mid+1, right)
		merge(arr, temp, left, mid, right)
	}
}

func merge(arr, temp []int, left, mid, right int) {
	i := left
	j := mid + 1
	t := 0
	for i < mid && j < left {
		if arr[i] < arr[j] {
			temp[t] = arr[i]
			t++
			i++
		} else {
			temp[t] = arr[j]
			t++
			j++
		}
	}
	for i <= mid {
		temp[t] = arr[i]
		t++
		i++
	}
	for j <= right {
		temp[t] = arr[j]
		t++
		j++
	}
	t = 0
	for left <= right {
		arr[left] = temp[t]
		left++
		t++
	}
}
