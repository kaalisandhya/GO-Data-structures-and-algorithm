1.package main

import (
	"fmt"
	"os"
	"strconv"
)


func insertionSort(array []float64) {
	for i := 0; i < len(array); i++ {
		temp := array[i]
		j := i - 1
		for ; j >= 0 && array[j] > temp; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = temp
	}
}

func bucketSort(array []float64, bucketSize int) []float64 {
	var max, min float64
	for _, n := range array {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	nBuckets := int(max-min)/bucketSize + 1
	buckets := make([][]float64, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]float64, 0)
	}

	for _, n := range array {
		idx := int(n-min) / bucketSize
		buckets[idx] = append(buckets[idx], n)
	}

	sorted := make([]float64, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			insertionSort(bucket)
			sorted = append(sorted, bucket...)
		}
	}

	return sorted
}

func main() {

	array :=[]float64{4, 10, 1, 3, 2, 2, 7, -6, 0,98,67,7,7}
	for _, arg := range os.Args[1:] {
		if n, err := strconv.ParseFloat(arg, 64); err == nil {
			array = append(array, n)
		}
	}
	fmt.Printf("Before Bucket Sort %v\n", array)
	array = bucketSort(array, 5)
	fmt.Printf("After Bucket Sort %v\n", array)
}

Before Bucket Sort [4 10 1 3 2 2 7 -6 0 98 67 7 7]
After Bucket Sort [-6 0 1 2 2 3 4 7 7 7 10 67 98]

Program exited.
