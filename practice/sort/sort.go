package sort

import "sort"

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xau4ci/
func topKFrequent(nums []int, k int) []int {
	mapp := make(map[int]int)
	for i := range nums {
		mapp[nums[i]]++
	}
	nl := make(nodeList, 0, len(mapp))
	for k, v := range mapp {
		nl = append(nl, node{
			key:   k,
			count: v,
		})
	}
	//按照频数由大到小排序
	sort.Sort(nl)
	var ret []int

	for i := 0; i < k; i++ {
		ret = append(ret, nl[i].key)
	}
	return ret
}

type node struct {
	key   int //key
	count int //频数
}

//实现sort, 按照频数由大到小排序
type nodeList []node

func (n nodeList) Len() int {
	return len(n)
}

func (n nodeList) Less(i, j int) bool {
	return n[i].count > n[j].count
}

func (n nodeList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

// https://leetcode-cn.com/problems/wiggle-sort-ii/
// 摆动排序2
//先排序，再逆序穿插
//逆序穿插而不是正序穿插是因为逆序能把前半段和后半段相同的分开来
//逆序：[1,2,2,3]排序后是[1,2,2,3],前半段是[1,2]，后半段是[2,3],穿插后是[2,3,1,2]
//正序：[1,2,2,3]排序后是[1,2,2,3],前半段是[1,2]，后半段是[2,3],穿插后是[1,2,2,3](不满足要求)
func wiggleSort(nums []int) {
	length := len(nums)
	tempNums := make([]int, length)
	copy(tempNums, nums)
	sort.Ints(tempNums)                 //从小到大排序
	idx1, idx2 := (length+1)>>1, length //idx1:指向前半段数组的尾部，inx2指向后半段数组的尾部
	for i := 0; i < length; i++ {
		if i&1 == 1 { //偶数位
			idx2--
			nums[i] = tempNums[idx2]
		} else { //奇数位,基数时，前半段会多一个，所以必须要先插前半段
			idx1--
			nums[i] = tempNums[idx1]
		}
	}
}
