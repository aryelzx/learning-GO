/*
https://leetcode.com/problems/two-sum/
*/
package main

import "fmt"

func main() {
	numsParams := []int{2, 7, 11, 15}
	target := 9
	a := twoSum(numsParams, target)
	fmt.Println(a)
}

var (
	firstSum  int
	secondSum int
)

func twoSum(nums []int, target int) []int {
	arr := []int{}
	result := []int{}

	for i, _ := range nums {
		arr = append(nums)
		for j, _ := range nums {
			prox := i + 1
			count := arr[j] + arr[prox]
			if count == target {
				firstSum = arr[j]
				secondSum = arr[prox]
				fmt.Println(firstSum, secondSum)
				// result = append(arr, arr[j], arr[prox])
			}
		}
		return arr
	}
	return result
}
