package array

import (
	"math"

	my_math "github.com/GGGxie/dataStructure/pkg/math"
)

type Array struct {
	elements []interface{}
	size     int
}

// 判断数组中的元素是否唯一
// true:唯一
// false:不唯一
func (a *Array) Isunique() bool {
	mapp := make(map[interface{}]int)
	for _, ch := range a.elements {
		mapp[ch]++
		if mapp[ch] > 1 {
			return false
		}
	}
	return true
}

// 给定一个整数数组，找出总和最大的连续数列，并返回总和。
// 动态规划
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 数组列表中的最大距离
// https://leetcode.cn/problems/maximum-distance-in-arrays/
func maxDistance(arrays [][]int) int {
	if len(arrays) == 0 {
		return 0
	}
	res, minVal, maxVal := math.MinInt, arrays[0][0], arrays[0][len(arrays[0])-1]
	for i := 1; i < len(arrays); i++ {
		res = my_math.Max(res, my_math.Abs(arrays[i][len(arrays[i])-1]-minVal))
		res = my_math.Max(res, my_math.Abs(maxVal-arrays[i][0]))
		minVal = my_math.Min(minVal, arrays[i][0])
		maxVal = my_math.Max(maxVal, arrays[i][len(arrays[i])-1])
	}
	return res
}

// 易混淆数,判断原数字旋转 180° 以后是否可以得到新的数字。
// https://leetcode.cn/problems/confusing-number/
func confusingNumber(n int) bool {
	// 记录翻转后的数字
	mapp := map[int]int{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}
	// 位处理
	var changeNum int

	for tmp := n; tmp != 0; tmp = tmp / 10 {
		z := tmp % 10
		if c, ok := mapp[z]; ok {
			changeNum = changeNum*10 + c
		} else {
			return false
		}
	}
	return changeNum != n
}

// 字符串的左右移
// https://leetcode.cn/problems/perform-string-shifts/?envType=study-plan-v2&envId=premium-algo-100
func stringShift(s string, shift [][]int) string {
	// 计算移动结果
	var mov int
	for i := range shift {
		// 左移+,右移-
		if shift[i][0] == 0 {
			mov += shift[i][1]
		} else {
			mov -= shift[i][1]
		}
	}
	// 计算首字母下标
	length := len(s)
	var idx int
	if mov > 0 { //向右移,取余
		idx = mov % length
	} else { //向左移,先取余再加总长度
		idx = (mov % length) + length
	}
	return s[idx:] + s[:idx]
}

// 相隔为 1 的编辑距离
// https://leetcode.cn/problems/one-edit-distance/
func isOneEditDistance(s string, t string) bool {
	distance := len(s) - len(t)
	if distance == 1 && (len(s) == 0 || len(t) == 0) {
		return true
	}
	var count int
	switch distance { //distance有三种情况
	case 0:
		{ //长度相等
			for i := range s {
				if s[i] != t[i] {
					count++
				}
			}
			return count == 1
		}
	case 1:
		{ //s 比 t 多一个
			for i, j := 0, 0; i < len(s) && j < len(t); i, j = i+1, j+1 {
				if s[i] != t[j] { //遇到不同的直接比较后半段
					count++
					return s[i+1:] == t[j:]
				}
			}
			return true //全部遍历完说明前面都相同,就最后一个字符不同,返回 true
		}
	case -1:
		{ //s 比 t 少一个
			for i, j := 0, 0; i < len(t) && j < len(s); i, j = i+1, j+1 {
				if s[j] != t[i] { //遇到不同的直接比较后半段
					count++
					return t[i+1:] == s[j:]
				}
			}
			return true //全部遍历完说明前面都相同,就最后一个字符不同,返回 true
		}
	default: //长度相差>1
		return false
	}
}
