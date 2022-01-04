package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := []int{1, 6, 3, 5}
	fmt.Println(exchange(a))
}

// https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
// 调整数组顺序使奇数位于偶数前面
// 双指针法
func exchange(nums []int) []int {
	length := len(nums)
	idx1, idx2 := 0, length-1 //idx1指向奇数，idx2指向偶数
	for {
		for { //for循环找到奇数
			if idx1 >= length || nums[idx1]&1 == 0 {
				break
			}
			idx1++
		}
		for { //for循环找到偶数
			if idx2 < 0 || nums[idx2]&1 == 1 {
				break
			}
			idx2--
		}
		if idx1 >= length || idx2 < 0 || idx1 >= idx2 { //下标判断
			break
		}
		nums[idx1], nums[idx2] = nums[idx2], nums[idx1] //交换
	}
	return nums
}
