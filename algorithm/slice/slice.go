package slice

// 找出所有行中最小公共元素
// https://leetcode.cn/problems/find-smallest-common-element-in-all-rows/
// 不能用 map,因为 map 是无序的,不能保证先遍历到最小公共元素,例如:
// mat = [[1,2,3],[2,3,4],[2,3,5]]
// 如果先遍历到 3 就错了
func smallestCommonElement(mat [][]int) int {
	slice := [10001]int{}
	length := len(mat)
	for _, s := range mat {
		for j := range s {
			slice[s[j]]++
		}
	}
	for i := range slice {
		if slice[i] >= length {
			return i
		}
	}

	return -1
}
