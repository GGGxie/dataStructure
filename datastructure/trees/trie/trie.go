package trie

// Test
// tree := trie.NewTrie([]string{"her", "xa"})
// flag := tree.Search("her")
// fmt.Println(flag)
// flag = tree.Search("xa1")
// fmt.Println(flag)
// flag = tree.StartsWith("h")
// fmt.Println(flag)
// flag = tree.StartsWith("ha")
// fmt.Println(flag)

//字典树
//作用是高效检索单词
// 常见应用场景
// 1.字符串检索
// 2.字符串最长公共前缀
// 3.排序
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

//判断字典树中是否存在该前缀的单词
func (t *TrieTree) StartsWith(prefix string) bool {
	nowNode := t.Root
	for _, v := range prefix {
		if nowNode.Next[v] == nil {
			return false
		}
		nowNode = nowNode.Next[v]
	}
	return true
}
