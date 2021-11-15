package slice

import (
	"fmt"
	"math"
	"math/rand"
)

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

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xmz79t/
// 多数元素。给定一个大小为 n 的数组，找到其中的多数元素。
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mapp := make(map[int]int)
	for i := range nums {
		mapp[nums[i]]++
	}
	maxNum := math.MinInt32
	index := 0
	for i := range mapp {
		if mapp[i] > maxNum {
			maxNum = mapp[i]
			index = i
		}
	}
	return index
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xm42hs/
// 旋转数组,给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
func rotate(nums []int, k int) {
	length := len(nums)
	k %= length //当k>length
	index := length - k
	copy(nums, append(nums[index:], nums[0:index]...))
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xm1rfd/
// 存在重复元素.给定一个整数数组，判断是否存在重复元素.
func containsDuplicate(nums []int) bool {
	mapp := make(map[int]int)
	for i := range nums {
		mapp[nums[i]]++
		if mapp[nums[i]] > 1 {
			return true
		}
	}
	return false
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xmy9jh/
// 移动零.给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 双指针实现,
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums) //left:第一个为0的下标,rigth:遍历切片的下标
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xmcbym/
// 两个数组的交集 II.给定两个数组，编写一个函数来计算它们的交集。(数组中可能存在重复数据)
func intersect(nums1 []int, nums2 []int) []int {
	//队列从小到大排序
	quickSort(nums1)
	quickSort(nums2)
	var ret []int
	//i:nums1下标,j:nums2下标
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			ret = append(ret, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		}
	}
	return ret
}

func quickSort(slice []int) {
	var (
		_quickSort func(left, right int, slice []int)     //利用递归不断对分区进行排序
		partition  func(left, right int, slice []int) int //排序
	)
	partition = func(left, right int, slice []int) int {
		flag := left      //基准
		index := left + 1 //标记比slice[flag]大的位置
		for i := index; i <= right; i++ {
			if slice[i] < slice[flag] {
				slice[i], slice[index] = slice[index], slice[i]
				index++
			}
		}
		slice[flag], slice[index-1] = slice[index-1], slice[flag]
		return (index - 1)
	}
	_quickSort = func(left, right int, slice []int) {
		if left < right {
			partitionIndex := partition(left, right, slice) //排序并获取基准位置
			//以基准位置进行分区，进行再排序
			_quickSort(left, partitionIndex-1, slice)
			_quickSort(partitionIndex+1, right, slice)
		}
	}
	left, right := 0, len(slice)-1 //left起始值下标，right末尾值下标
	_quickSort(left, right, slice)
}

// https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
// 搜索二维矩阵 II,
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
// 暴力对每一行进行二分
func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		if binarySearch(row, target) {
			return true
		}
	}
	return false
}

// https://leetcode-cn.com/problems/product-of-array-except-self/solution/chu-zi-shen-yi-wai-shu-zu-de-cheng-ji-by-leetcode-/
// 除自身以外数组的乘积
//所有非0元素相乘得到sum
//再用sum/nums[i]得到返回切片为i的值
func productExceptSelf(nums []int) []int {
	sum := 1     //非0总乘积
	numZero := 0 //切片中0的个数
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			numZero++
			continue
		}
		sum *= nums[i]
	}
	for i := range nums {
		if numZero > 1 { //多0切片
			nums[i] = 0
		} else if numZero > 0 { //单0切片
			if nums[i] == 0 && numZero != len(nums) { //预防{0}
				nums[i] = sum
			} else {
				nums[i] = 0
			}
		} else { //无0切片
			nums[i] = sum / nums[i]
		}
	}
	return nums
}

// 官方实现
// func productExceptSelf(nums []int) []int {
// 	length := len(nums)
// 	answer := make([]int, length)
// 	// answer[i] 表示索引 i 左侧所有元素的乘积
// 	// 因为索引为 '0' 的元素左侧没有元素， 所以 answer[0] = 1
// 	answer[0] = 1
// 	for i := 1; i < length; i++ {
// 		answer[i] = nums[i-1] * answer[i-1]
// 	}
// 	// R 为右侧所有元素的乘积
// 	// 刚开始右边没有元素，所以 R = 1
// 	R := 1
// 	for i := length - 1; i >= 0; i-- {
// 		// 对于索引 i，左边的乘积为 answer[i]，右边的乘积为 R
// 		answer[i] = answer[i] * R
// 		// R 需要包含右边所有的乘积，所以计算下一个结果时需要将当前值乘到 R 上
// 		R *= nums[i]
// 	}
// 	return answer
// }

//二分搜索
func binarySearch(list []int, target int) bool {
	if len(list) == 0 {
		return false
	}
	mid := len(list) / 2
	if list[mid] == target {
		return true
	}
	if list[mid] > target {
		return binarySearch(list[0:mid], target) //当len(list)=1,list[0:0]为空，len(list[0:0])=0
	} else if list[mid] < target {
		return binarySearch(list[mid+1:], target) //当len(list)=1,list[1:]为空,len(list[1:])=0
	}
	return false
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://leetcode-cn.com/problems/basic-calculator-ii/
// 基本计算器 II
// 解题思路：用数组模拟栈
func calculate(s string) int {
	var tempSlice []int //临时数组，记录所有待相加元素
	var num int
	preSign := '+'
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit { //记录数字
			num = num*10 + int(ch-'0')
		}
		if (!isDigit && ch != ' ') || i == len(s)-1 { //ch为符号 || 遍历到最后一个字符
			switch preSign {
			case '+': //直接加入数组
				tempSlice = append(tempSlice, num)
			case '-': //将负数加入数组
				tempSlice = append(tempSlice, -num)
			case '*': // 将数组最后一个拿出来*num，将结果压入数组
				tempSlice[len(tempSlice)-1] *= num
			case '/': // 将数组最后一个拿出来/num，将结果压入数组
				tempSlice[len(tempSlice)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	var sum int
	for i := range tempSlice { //将临时数组所有元素相加
		sum += tempSlice[i]
	}
	return sum
}

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xa6dkt/
// 26进制转10进制
func titleToNumber(columnTitle string) int {
	var ret int
	length := len(columnTitle)
	for i, ch := range columnTitle {
		temp := exponent(26, (length - i - 1)) //位数
		temp2 := (int(ch-'A') + 1)             //最高位
		ret += temp2 * temp
	}
	return ret
}

//计算a的n次方
func exponent(a, n int) int {
	result := int(1)
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= a
		}
		a *= a
	}
	return result
}

// https://leetcode-cn.com/problems/insert-delete-getrandom-o1/
// O(1) 时间插入、删除和获取随机元素
// 通过把要删除元素和数组尾部元素调换，然后去掉尾部元素，实现O(1)删除
type RandomizedSet struct {
	mapp map[int]int //记录数据的下标
	data []int       //记录数据的数组
}

func RandomizedSetConstructor() RandomizedSet {
	return RandomizedSet{
		mapp: make(map[int]int),
		data: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.mapp[val]; !ok {
		this.mapp[val] = len(this.data)
		this.data = append(this.data, val)
		return true
	}
	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.mapp[val]; ok {
		length := len(this.data)
		this.data[index], this.data[length-1] = this.data[length-1], this.data[index] //数组尾部和下标互换
		this.data = this.data[0 : length-1]                                           //去掉尾部
		if index != length-1 {                                                        //如果刚好移除的是尾部元素，就不用替换下标
			this.mapp[this.data[index]] = index //被替换的尾部元素改下标
		}
		delete(this.mapp, val) //去掉下标存储
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	length := len(this.data)
	random := rand.Int() % length
	return this.data[random]
}

// https://leetcode-cn.com/problems/shuffle-an-array/
// 打乱数组,洗牌算法
type Solution struct {
	nums []int
}

func SolutionConstructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	ret := make([]int, len(this.nums))
	copy(ret, this.nums)
	return ret
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int { //模拟洗牌算法
	ret := make([]int, len(this.nums))
	copy(ret, this.nums)
	for i := range ret {
		random := rand.Intn(i + 1) //返回一个[0,i+1)的随机值
		ret[i], ret[random] = ret[random], ret[i]
	}
	return ret
}

// https://leetcode-cn.com/problems/gas-station/
// 加油站
// 模拟题,遍历数组
func canCompleteCircuit(gas []int, cost []int) int {
	length := len(cost)
	tempList := make([]int, length)    //标记从i节点开始，到最后因为缺少多少汽油而停止
	for i := length - 1; i >= 0; i-- { //从最后一个节点开始遍历，然后往前去遍历每一个节点开始的情况（因为这样子后面节点的tempList都有值了）
		if gas[i] < cost[i] {
			continue
		}
		tempGas := 0 //初始汽油值
		count := 0
		for j := i; ; j = ((j + 1) % length) { //从i节点开始往前遍历
			count++
			tempGas += gas[j]
			tempGas -= cost[j]
			if tempGas < 0 { //如果从i点开始，到了j点，汽油不够了，就记录到tempList中
				tempList[i] = tempGas
				break
			}
			if tempGas < tempList[j] { //如果从i点开始，到了j点，汽油还充足，但是少于从j开始到最后缺少的汽油，则没必要继续遍历
				tempList[i] = tempGas - tempList[j]
				break
			}
			if count == length {
				fmt.Println(tempList)
				return i
			}
		}

	}
	return -1
}
