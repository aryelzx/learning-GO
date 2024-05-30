/*
https://leetcode.com/problems/two-sum/
*/
package main

import "fmt"

func main() {
	numsParams := []int{3, 2, 4}
	target := 6
	a := twoSum(numsParams, target)
	fmt.Println(a, "two sums result")
	indices := findIndex(numsParams, a)
	fmt.Println(indices, "indices")
}

func twoSum(nums []int, target int) []int {
	arr := nums
	result := []int{}
	for i, _ := range nums {
		for j, _ := range arr {
			prox := i + 1
			count := arr[j] + arr[prox]
			firstSum := arr[j]
			secondSum := arr[prox]
			if count == target {
				result = append(result, firstSum, secondSum)
				return result
			}

		}
	}
	return result
}

func findIndex(slice []int, value []int) []int {
	indices := []int{}
	for _, v := range slice {
		for i, p := range value {
			if p == v {
				indices = append(indices, i)
			}
		}
	}
	return indices
}
