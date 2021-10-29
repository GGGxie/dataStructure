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
