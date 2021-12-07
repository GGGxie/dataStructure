package dp

import (
	"math"
	"strings"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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

// https://leetcode-cn.com/leetbook/read/top-interview-questions/x2959v/
// 完全平方数:需要让组成和的完全平方数的个数最少,给你一个整数 n ，返回和为 n 的完全平方数的最少数量
func numSquares(n int) int {
	dp := make([]int, n+1) //dp[i]：达到i所需要的最少个数
	dp[0] = 0
	for i := 1; i <= n; i++ {
		minNum := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minNum = min(minNum, dp[i-j*j])
		}
		dp[i] = minNum + 1
	}
	return dp[n]
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/x29fxj/
// 最长上升子序列
// 状态转移方程：dp[i]=max(dp[j])+1,其中0≤j<i且num[j]<num[i]
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums)) //dp[i]:到达i的最长上升子序列的个数
	maxNum := math.MinInt32
	for i := 0; i < len(nums); i++ {
		count := 1
		for j := i; j >= 0; j-- {
			if nums[j] < nums[i] {
				count = max(dp[j]+1, count)
			}
		}
		dp[i] = count
		if maxNum < dp[i] {
			maxNum = dp[i]
		}
	}
	return maxNum
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/x2echt/
// 零钱兑换，最少的硬币个数 。如果没有任何一种硬币组合能组成总金额
// 动态转移方程 dp[i]=min(dp[i-coins[j]]+1),其中coins[j] <= i
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) //dp[i]：i代表amount为i时，需要的最少硬币数
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		minNum := math.MaxInt32
		for j := range coins {
			if coins[j] <= i {
				minNum = min(minNum, dp[i-coins[j]]+1)
			}
		}
		dp[i] = minNum
	}
	if dp[amount] == math.MaxInt32 { //没有匹配返回-1
		return -1
	} else {
		return dp[amount]
	}
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xafdmc/
// 至少有K个重复字符的最长子串，给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。
// 分治实现，找出s中数量小于k的字母，然后根据字母来分割字符串，再从各个字符串中找出最长子字符串。
func longestSubstring(s string, k int) int {
	mapp := make(map[int]int) //记录s中每种字符的数量
	for _, r := range s {
		index := int(r - 'a')
		mapp[index]++
	}
	var temp string //记录数量小于k的字母
	for i, _ := range mapp {
		if mapp[i] < k {
			temp = string(i + 'a')
			break
		}
	}
	if temp == "" {
		return len(s)
	}
	maxNum := math.MinInt32
	for _, child := range strings.Split(s, string(temp)) {
		maxNum = max(longestSubstring(child, k), maxNum)
	}

	//记录s中
	return maxNum
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xmk3rv/
// 乘积最大子数组
// 动态转移方程：
// maxDp[i] = max{maxDp[i-1]*nums[i], minDp[i-1]*nums[i], nums[i]}
// minDp[i] = min(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]), nums[i]
func maxProduct(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	maxDp := make([]int, length) //maxDP[i]:以第i个元素结尾的最大子数组的乘积
	minDp := make([]int, length) //minDP[i]:以第i个元素结尾的最小子数组的乘积
	maxDp[0], minDp[0] = nums[0], nums[0]
	maxNum := nums[0]
	for i := 1; i < length; i++ {
		maxDp[i] = max(max(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]), nums[i])
		minDp[i] = min(min(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]), nums[i])
		maxNum = max(max(maxDp[i], minDp[i]), maxNum)
	}
	return maxNum
}

// https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/solution/go-gun-dong-shu-zu-by-xilepeng-x8wj
// 青蛙跳台阶问题
// 动态规划转移方程numWays(n)=numWays(n-1)+numWays(n-2)
func numWays(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	dp := make([]int, n+1) //记录跳到i级台阶有几种方法
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007 //加法分配律？
	}
	return dp[n]
}

// // //优化方案，空间为O(1)
// func numWays(n int) int {
// 	prev, curr := 1, 1
// 	for i := 2; i <= n; i++ {
// 		next := (prev + curr) % 1000000007
// 		prev = curr
// 		curr = next
// 	}
// 	return curr
// }

// https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof
// 连续子数组的最大和
// 动态规划: maxSubArray(n)=max(maxSubArray(n-1)+nums[n],nums[n])
// 用滚动数组优化空间
func maxSubArray(nums []int) int {
	temp, ret := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		temp2 := max(nums[i], temp+nums[i]) //temp:记录以nums(i-1)结束的最大和,temp2:记录nums(i)结束的最大和
		if temp2 > ret {
			ret = temp2
		}
		temp = temp2 //滚动,temp记录nums(i)结束的最大和,继续遍历i+1
	}
	return ret
}
