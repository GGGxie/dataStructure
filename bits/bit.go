package bits

import (
	"fmt"
)

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

//翻转数位,你可以将一个数位从0变为1,找出你能够获得的最长的一串1的长度。
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
