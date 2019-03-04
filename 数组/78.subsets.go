package main

import (
	"fmt"
	"math"
)

//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
//
//说明：解集不能包含重复的子集。
//
//示例:
//
//输入: nums = [1,2,3]
//输出:
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]
func main() {
	test := []int{}
	fmt.Println(subsets(test))

}
func subsets(nums []int) [][]int {
	result := make([][]int, 0, 1)
	if len(nums) == 0 {
		pre := []int{}
		result = append(result, pre)
		return result
	}
	len := float64(len(nums))
	co := int(math.Pow(2.0, len))
	result = make([][]int, 0, co)
	// 思路：nums的长度，二进制，0，1。
	n := 0
	for n < co {
		subset := []int{}
		index := 0
		temp := n
		for temp > 0 {

			if temp&1 == 1 {
				subset = append(subset, nums[index])
			}
			index += 1
			temp = temp >> 1
		}
		n += 1
		result = append(result, subset)
	}
	return result
}
