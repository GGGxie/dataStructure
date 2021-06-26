package bits

import (
	"fmt"
)

//常用操作
//取出最地位：num & 1
//取出后n位：num & ((1<<n)-1)

//用M去填充N的第i-j位
// 输入：N = 1024(10000000000), M = 19(10011), i = 2, j = 6
// 输出：N = 1100(10001001100)
func InsertBits(N int, M int, i int, j int) int {
	for k := i; k <= j; k++ { //把N的第i-j位置为0
		N &= (^(1 << k))
		fmt.Printf("%016b\n", N)
	}
	fmt.Printf("%016b\n", N)
	N |= (M << i) //把N的第i-j位和M进行|运算
	return N
}

//输出十进制小数的二进制表示
func PrintBin(num float64) string {
	var ret []byte
	var count int
	ret = append(ret, []byte{'0', '.'}...)
	for ; count < 32; count++ {
		temp := num * 2
		if temp >= 1 {
			ret = append(ret, '1')
			num = temp - 1
		} else {
			ret = append(ret, '0')
			num = temp
		}
		if num == 0 {
			break
		}
	}

	if count < 32 {
		return string(ret)
	} else {
		return "ERROR"
	}
}

//翻转数位,你可以将一个数位从0变为1,找出你能够获得的最长的一串1的长度。TODO:改成动态规划
func reverseBits(num int) int {
	//pre记录从上一个0开始，1的总数
	//max记录最大的1的总数
	//current记录用0连接的两段的1的总数
	var pre, max, current int
	var flag bool
	for i := 0; i < 32; i++ { //遍历num
		a := num & 1 //取出个位
		num = num >> 1
		if a == 1 {
			current++
			pre++
		} else if a == 0 {
			if !flag { //替换1还不被使用
				current++
				flag = true
			} else { //替换1已被使用
				current = pre + 1
				if current > max {
					max = current
				}
			}
			pre = 0
		}
		if current > max {
			max = current
		}
	}
	return max
}

// 给定一个正整数，找出与其二进制表达式中1的个数相同且大小最接近的那两个数（一个略大，一个略小）。
// 解题思路
// 比 num 大的数：从右往左，找到第一个 01 位置，然后把 01 转为 10，右侧剩下的 1 移到右侧的低位，右侧剩下的位清0。
// 比 num 小的数：从右往左，找到第一个 10 位置，然后把 10 转为 01，右侧剩下的 1 移到右侧的高位，右侧剩下的位置0。
func findClosedNumbers(num int) []int {
	larger, smaller := -1, -1
	count := func(num int) int {
		var sum int
		for num != 0 {
			if num&1 == 1 {
				sum++
			}
			num >>= 1
		}
		return sum
	}
	sumOfNum := count(num)
	//获取偏大值
	temp := num + 1
	for temp <= 2147483647 {
		if count(temp) == sumOfNum {
			larger = temp
			break
		}
		temp++
	}

	//获取偏小值
	temp = num - 1
	for temp >= 1 {
		if count(temp) == sumOfNum {
			smaller = temp
			break
		}
		temp--
	}
	return []int{larger, smaller}
}

//整数转换。编写一个函数，确定需要改变几个位才能将整数A转成整数B。
//todo:为什么获取数字的二进制表示中1的个数要这样写
//todo:为什么要int32?
func convertInteger(A int, B int) int {
	count := func(num int32) (sum int) { //获取数字的二进制表示中1的个数
		for num != 0 {
			sum++
			num = num & (num - 1)
		}
		return
	}
	return count(int32(A ^ B))
}

// 编写程序，交换某个整数的奇数位和偶数位
func exchangeBits(num int) int {
	//取出奇数位
	single := ((num & 0x55555555) << 1)
	//取出偶数位
	double := ((num & 0xaaaaaaaa) >> 1)
	//基数位和偶数位进行或操作
	return single | double
}

// 05.08. 绘制直线
func drawLine(length int, w int, x1 int, x2 int, y int) []int {
	ret := make([]int, length)
	for i := range ret { //遍历元素
		var temp int32                        //TODO:为什么要int32不能用int
		for z := 0 + 32*i; z < 32+32*i; z++ { //遍历元素的bit位
			if z >= x1+y*w && z <= x2+y*w { //判断是否在x1和x2之间
				temp |= 1 << (31 - (z - 32*i)) //在x1和x2之间的0全部置为1
			}
		}
		ret[i] = int(temp)
	}
	return ret
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
