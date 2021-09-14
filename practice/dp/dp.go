package dp

import (
	"math"
)

// 三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。
func waysToStep(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	default:
		dp1 := 1
		dp2 := 2
		dp3 := 4
		for i := 4; i <= n; i++ {
			dp1, dp2, dp3 = dp2, dp3, (dp1+dp2+dp3)%1000000007 //每次都要取余？？？？？
		}
		return dp3
	}
}

// 硬币。给定数量不限的硬币，币值为25分、10分、5分和1分，编写代码计算n分有几种表示法。(结果可能会很大，你需要将结果模上1000000007)
func waysToChange(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	coins := []int{1, 5, 10, 25}
	for i := 0; i < 4; i++ {
		for j := 1; j <= n; j++ {
			if j-coins[i] >= 0 {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[n] % 1000000007
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/x2hnpi/
// 二叉树中的最大路径和
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于0时，才会选取对应子节点
		// 如果两边贡献值都小于0,那么两边都可以不要
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		priceNewPath := node.Val + leftGain + rightGain
		// 更新答案
		maxSum = max(maxSum, priceNewPath)
		// 返回节点的最大贡献值
		return max(node.Val+leftGain, node.Val+rightGain)
	}
	maxGain(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/x2xmre/
// 最长连续序列
// hash表实现,时间复杂度O(n),空间复杂度O(n)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mapp := make(map[int]bool)
	max := math.MinInt32
	for i := range nums {
		mapp[nums[i]] = true
	}
	for i := range nums {
		if !mapp[nums[i]-1] { //起点
			count := 1
			start := nums[i]
			for mapp[start+1] {
				count++
				start += 1
			}
			if count > max {
				max = count
			}
		}
	}
	return max
}
