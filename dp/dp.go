package dp

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
