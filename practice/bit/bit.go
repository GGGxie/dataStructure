package bit

// https://leetcode-cn.com/problems/single-number/
// 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 思路:数组所有元素进行异或即可
// 1. 0^A=A
// 2. A^A=A
// 3. 异或运算满足交换律和结合律  a^b^a=b^a^a=b^(a^a)=b^0=b。
func singleNumber(nums []int) int {
	ret := 0
	for i := range nums {
		ret ^= nums[i]
	}
	return ret
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xaph0j/
//找到第一个不重复的字符
//用hash来存储频数
func firstUniqChar(s string) int {
	mapp := make(map[rune]int)
	for _, r := range s {
		mapp[r]++
	}
	for i, r := range s {
		if mapp[r] <= 1 {
			return i
		}
	}
	return -1
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xar9lv/
// 字母异位词
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapp := make(map[rune]int)
	for _, r := range s {
		mapp[r]++
	}
	for _, r := range t {
		mapp[r]--
	}
	for i := range mapp {
		if mapp[i] != 0 {
			return false
		}
	}
	return true
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/x2dx36/
// 颠倒二进制位
func reverseBits(num uint32) uint32 {
	var ret uint32
	for i := 0; i < 32; i++ {
		temp := num & 1 //取出num最后一个二进制位
		num >>= 1       //num右移一位
		ret <<= 1       //ret左移一位
		ret += temp     //ret+取出的二进制位
	}
	return ret
}
