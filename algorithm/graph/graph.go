package graph

// 孤独像素 I
// https://leetcode.cn/problems/lonely-pixel-i/
// 先进行预处理,统计各行各列存在'B'的数量,最后遍历的时候为行列为 1 的都是孤独像素
func findLonelyPixel(picture [][]byte) int {
	rowMap := map[int]int{}
	colMap := map[int]int{}
	for i := range picture {
		for j := range picture[i] {
			if picture[i][j] == 'B' {
				rowMap[i]++
				colMap[j]++
			}
		}
	}
	ret := 0
	for i := range picture {
		for j := range picture[i] {
			if picture[i][j] == 'B' && rowMap[i] == 1 && colMap[j] == 1 {
				ret++
			}
		}
	}
	return ret
}

// 有效的单词方块
// https://leetcode.cn/problems/valid-word-square/
// 直接模拟
func validWordSquare(words []string) bool {
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words) || i >= len(words[j]) || words[i][j] != words[j][i] {
				return false
			}
		}
	}
	return true
}
