package math

import (
	"math"
	"strconv"
)

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

// https://leetcode-cn.com/problems/happy-number/
// 快乐数
func isHappy(n int) bool {
	temp := n
	mapp := make(map[int]bool) //标记环
	for {
		var tempB int //计算平均和的值
		for temp != 0 {
			z := temp % 10
			temp /= 10
			tempB += z * z
		}
		if mapp[tempB] { //已被遍历，进入循环
			break
		}
		if tempB == 1 {
			return true
		}
		mapp[tempB] = true
		temp = tempB
	}
	return false
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xm6kpg/
// Fizz Buzz
// 模拟题
func fizzBuzz(n int) []string {
	ret := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ret = append(ret, "FizzBuzz")
		} else if i%3 == 0 {
			ret = append(ret, "Fizz")
		} else if i%5 == 0 {
			ret = append(ret, "Buzz")
		} else {
			ret = append(ret, strconv.FormatInt(int64(i), 10))
		}
	}
	return ret
}

// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
// 斐波那契数列
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	temp1 := 0 //指向f(n-2)
	temp2 := 1 //指向f(n-1)
	for i := 2; i <= n; i++ {
		temp3 := (temp1 + temp2) % 1000000007 //计算出f(n)
		temp1 = temp2                         //f(n-2)指向之前f(n-1)的位置
		temp2 = temp3                         //f(n-1)指向之前f(n)的位置
	}
	return temp2
}
