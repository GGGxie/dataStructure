package string

import "strings"

// https://leetcode-cn.com/leetbook/read/top-interview-questions/xah8k6/
// 判断字符串是否为回文串
func isPalindrome(s string) bool {
	stemp := strings.ToUpper(s) //全部转大写
	length := len(s)
	for left, right := 0, length-1; left < right; left, right = left+1, right-1 {
		for !isLetterOrNumber(stemp[left]) { //取下一个字母或数字进行比较
			left++
			if left > length-1 {
				break
			}
		}
		for !isLetterOrNumber(stemp[right]) { //取下一个字母或数字进行比较
			right--
			if right < 0 {
				break
			}
		}
		if left < right && stemp[left] != stemp[right] {
			return false
		}
	}
	return true
}

func isLetterOrNumber(ch byte) bool { //判断是否为字母或数字
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func wordBreak(s string, wordDict []string) bool {
	t := NewTrie(wordDict)
	t.Insert(s)
	return t.Search(s)
}

//树的root为空,不储存数据
type TrieTree struct {
	Root *TrieNode
}

//树节点
type TrieNode struct {
	IsWord bool               //记录到该节点是否为一个单词
	Next   map[rune]*TrieNode //记录下一个节点的信息
}

//新建字典树
func NewTrie(arrList []string) *TrieTree {
	trie := &TrieTree{}
	trie.Root = &TrieNode{
		Next:   make(map[rune]*TrieNode),
		IsWord: false,
	}
	//插入单词
	for _, value := range arrList {
		trie.Insert(value)
	}
	return trie
}

//添加新单词
func (t *TrieTree) Insert(word string) {
	nowNode := t.Root
	for _, v := range word {
		if nowNode.Next[v] == nil { //子节点为空则添加新节点
			newNode := &TrieNode{
				Next:   make(map[rune]*TrieNode),
				IsWord: false,
			}
			nowNode.Next[v] = newNode
		}
		//遍历下一个节点
		nowNode = nowNode.Next[v]
	}
	//标记到达该节点为一个单词
	nowNode.IsWord = true
}

//查找单词
func (t *TrieTree) Search(word string) bool {
	nowNode := t.Root
	for _, v := range word {
		if nowNode.Next[v] == nil { //遍历不下去则证明该单词不存在
			return false
		}
		nowNode = nowNode.Next[v]
	}
	//返回该节点是否一个单词
	return nowNode.IsWord
}
