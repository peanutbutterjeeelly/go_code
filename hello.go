package main

import(
	"fmt"
)

func twoSumLeetcode001(nums []int, target int) []int {
    for i, x := range nums {
        for j := i + 1; j < len(nums); j++ {
            if x+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}

func main(){
	fmt.Println("something...")

	leetcode001Array := []int{2,7,11,13}
	leetcode001Target := 9
	fmt.Println(twoSumLeetcode001(leetcode001Array, leetcode001Target))
}