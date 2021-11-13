package math

import "math"

// https://leetcode-cn.com/problems/max-points-on-a-line/
// 直线上最多的点数
// 暴力遍历，以某个点为起点，再和其他点连线，k相同则在同一条直线上
func maxPoints(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	max := math.MinInt64
	for i, point := range points {
		cnt := make(map[float64]int) //存相同k值的点个数
		for j := i + 1; j < len(points); j++ {
			if i == j {
				continue
			}
			point2 := points[j]
			var k float64
			if (point[0] - point2[0]) == 0 { //当两个点在x轴上，k不能为0，一定是无穷大，如果记录为0，那么会和y轴上的点的斜率冲突
				k = math.Inf(1)
			} else {
				k = float64(point[1]-point2[1]) / float64(point[0]-point2[0]) //计算k值
			}
			cnt[k]++
			if cnt[k] > max {
				max = cnt[k]
			}
		}
	}
	return max + 1
}

// https://leetcode-cn.com/problems/count-primes/
// 计算小于n的质数个数
// 常规筛选法
func countPrimes(n int) int {
	flag := make([]bool, n+1) //标记所有元素，非素数就为true
	var ret int
	for i := 2; i <= n; i++ {
		if flag[i] {
			continue
		}
		for j := i; j <= n; j += i { //素数的所有倍数标记为true
			flag[j] = true
		}
		ret++
	}
	return ret
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/x2lkle/
// 3的幂
// 除了暴力还有更优的解法。
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 { //从3^n一直到3^0
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return true
}
