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

// https://leetcode-cn.com/problems/reverse-integer/
// 整数反转
func reverse(x int) int {
	ret := 0
	for x != 0 {
		temp := x % 10
		x /= 10
		ret *= 10
		ret += temp
	}
	if ret > int(math.MaxInt32) || ret < int(-math.MaxInt32) {
		return 0
	}
	return ret
}
