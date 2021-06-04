package bits

import "fmt"

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
