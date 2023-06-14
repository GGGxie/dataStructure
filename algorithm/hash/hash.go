package hash

// 找出变位映射
// https://leetcode.cn/problems/find-anagram-mappings/
// 虽然 A B中可能存在相同队列,但是输出其中一种就好
func anagramMappings(nums1 []int, nums2 []int) []int {
	mapp := map[int]int{}
	for k, v := range nums2 {
		mapp[v] = k
	}
	ret := make([]int, 0, len(nums1))
	for i := range nums1 {
		ret = append(ret, mapp[nums1[i]])
	}
	return ret
}

// 回文排列
// https://leetcode.cn/problems/palindrome-permutation/
// 回文队列中的字符个数为奇数少于 2 个
func canPermutePalindrome(s string) bool {
	mapp := map[byte]int{}
	for i := range s {
		mapp[s[i]]++
	}
	var count int
	for _, v := range mapp {
		if v%2 != 0 {
			count++
		}
	}
	return count < 2
}

// 句子相似性
// https://leetcode.cn/problems/sentence-similarity/
// 注意 similarPairs 中有 1:n 的情况
func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
	// 记录映射关系
	mapp := map[string]bool{}
	for _, s := range similarPairs {
		str := s[0] + "/" + s[1]
		mapp[str] = true
	}
	for i, v1 := range sentence1 {
		v2 := sentence2[i]
		str1, str2 := v1+"/"+v2, v2+"/"+v1
		if v1 != v2 && !mapp[str1] && !mapp[str2] {
			return false
		}
	}
	return true
}

// 单行键盘
// https://leetcode.cn/problems/single-row-keyboard/
// map 记录所有键盘字母的下标,然后进行距离计算
func calculateTime(keyboard string, word string) int {
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	mapp := map[byte]int{}
	for i := range keyboard {
		mapp[keyboard[i]] = i
	}
	var count int
	var idx int
	for i := range word {
		count += abs(mapp[word[i]] - idx)
		idx = mapp[word[i]]
	}
	return count
}

// 1133. 最大唯一数
// https://leetcode.cn/problems/largest-unique-number/
// 数的范围是0-1000,可以考虑用 map,但是用数组存比较省空间
func largestUniqueNumber(nums []int) int {
	s := [1001]int{}
	for _, v := range nums {
		s[v]++
	}
	for i := 1000; i >= 0; i-- {
		if s[i] == 1 {
			return i
		}
	}
	return -1
}

// 数元素
// https://leetcode.cn/problems/counting-elements/
// 数的范围是0-1000,可以考虑用 map,但是用数组存比较省空间
func countElements(arr []int) int {
	var count int
	slice := [1001]int{}
	for i := range arr {
		slice[arr[i]]++
	}
	for i := 0; i < 1000; i++ {
		if slice[i+1] != 0 {
			count += slice[i]
		}
	}
	return count
}
