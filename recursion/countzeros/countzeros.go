package main

import "fmt"

func main() {
	array := []int{0, 0, 0, 1, 2, 3, 4, 5, 0} // 4
	fmt.Println(headRecursiveCountZeroes(array, len(array)))
	fmt.Println(tailRecursiveCountZeroes(array, len(array)))
}

func headRecursiveCountZeroes(a []int, size int) int {
	if size == 0 {
		return 0
	}
	count := headRecursiveCountZeroes(a, size-1)
	if a[size-1] == 0 {
		count++
	}
	return count
}

func tailRecursiveCountZeroes(a []int, size int) int {
	if size == 0 {
		return 0
	}
	count := 0
	if a[size-1] == 0 {
		count++
	}
	return tailRecursiveCountZeroes(a, size-1) + count
}
