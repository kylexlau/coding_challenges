// Given an array of integers nums and an integer target, return indices of the
// two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may
// not use the same element twice.

// You can return the answer in any order.

// Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?

package main
import "fmt"

func twoSum(nums []int, target int) []int {
	for j:=0; j < len(nums); j++ {
		for k:=j+1; k<len(nums); k++{
			if nums[j]+nums[k] == target {
				return []int{j,k}
			}
		}
	}
	return []int{0,0}
}

func main() {
	a := []int{2,7,11,15}
	t := 26
	fmt.Println(twoSum(a,t))
}
