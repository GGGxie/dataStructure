package main

import (
	"fmt"
	"sync"

	"github.com/GGGxie/dataStructure/graph"
)

func main() {
	g := graph.NewGraph()
	n1, n2, n3, n4, n5 := graph.Node{1}, graph.Node{2}, graph.Node{3}, graph.Node{4}, graph.Node{5}

	g.AddNode(&n1)
	g.AddNode(&n2)
	g.AddNode(&n3)
	g.AddNode(&n4)
	g.AddNode(&n5)

	g.AddEdge(&n1, &n2)
	g.AddEdge(&n1, &n5)
	g.AddEdge(&n2, &n3)
	g.AddEdge(&n2, &n4)
	g.AddEdge(&n2, &n5)
	g.AddEdge(&n3, &n4)
	g.AddEdge(&n4, &n5)
	fmt.Println(g.Edges)
	g.BFS(&n4, func(node *graph.Node) {
		fmt.Println(node)
	})
}

type Node struct {
	value int
}

type Graph struct {
	nodes []*Node          // 节点集
	edges map[Node][]*Node // 邻接表表示的无向图
	lock  sync.RWMutex     // 保证线程安全
}

type NodeQueue struct {
	nodes []Node
	lock  sync.RWMutex
}

// 实现 BFS 遍历
func (g *Graph) BFS(f func(node *Node)) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	// 初始化队列
	q := NewNodeQueue()
	// 取图的第一个节点入队列
	head := g.nodes[0]
	q.Enqueue(*head)
	// 标识节点是否已经被访问过
	visited := make(map[*Node]bool)
	visited[head] = true
	// 遍历所有节点直到队列为空
	for {
		if q.IsEmpty() {
			break
		}
		node := q.Dequeue()
		visited[node] = true
		nexts := g.edges[*node]
		// 将所有未访问过的邻接节点入队列
		for _, next := range nexts {
			// 如果节点已被访问过
			if visited[next] {
				continue
			}
			q.Enqueue(*next)
			visited[next] = true
		}
		// 对每个正在遍历的节点执行回调
		if f != nil {
			f(node)
		}
	}
}

// 生成节点队列
func NewNodeQueue() *NodeQueue {
	q := NodeQueue{}
	q.lock.Lock()
	defer q.lock.Unlock()
	q.nodes = []Node{}
	return &q
}

// 入队列
func (q *NodeQueue) Enqueue(n Node) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.nodes = append(q.nodes, n)
}

// 出队列
func (q *NodeQueue) Dequeue() *Node {
	q.lock.Lock()
	defer q.lock.Unlock()
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return &node
}

// 判空
func (q *NodeQueue) IsEmpty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return len(q.nodes) == 0
}

// func waysToChange(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	dp := make([]int, n+1)
// 	dp[0] = 1
// 	for i := 1; i <= n; i++ {
// 		if i >= 1 {
// 			dp[i] += dp[i-1]
// 		}
// 		if i >= 5 {
// 			dp[i] += dp[i-5]
// 		}
// 		if i >= 10 {
// 			dp[i] += dp[i-10]
// 		}
// 		if i >= 25 {
// 			dp[i] += dp[i-25]
// 		}
// 		fmt.Println(i, dp[i])
// 	}
// 	return dp[n]
// }

// func multiply(A int, B int) int {
// 	var re func(A, B *int, C int)
// 	re = func(A, B *int, C int) {
// 		if *B == 0 {
// 			return
// 		}
// 		*A += C
// 		*B -= 1
// 		re(A, B, C)
// 	}
// 	if A > B {
// 		B -= 1
// 		re(&A, &B, A)
// 		return A
// 	} else {
// 		A -= 1
// 		re(&B, &A, B)
// 		return B
// 	}
// }

// func findClosedNumbers(num int) []int {
// 	count := func(num int) int { //找出1的总数
// 		var sum int
// 		for num != 0 {
// 			if num&1 == 1 {
// 				sum++
// 			}
// 			num >>= 1
// 		}
// 		return sum
// 	}

// 	//找到最大的值
// 	larger, smaller := -1, -1
// 	sigA := 0b01
// 	sigB := 0b10
// 	var index int
// 	for index = 0; index <= 30; index++ { //从右往左找到第一个01的位置
// 		if num&(sigA<<index) == sigA {
// 			temp := (num & ((1 << index) - 1)) //获取
// 			num = ((sigB << index) | temp)
// 			break
// 		}
// 	}
// 	//获取偏大值
// 	for ; index >= 0; index-- {

// 	}
// 	//获取偏小值
// 	temp = num - 1
// 	for temp >= 1 {
// 		if count(temp) == sumOfNum {
// 			smaller = temp
// 			break
// 		}
// 		temp--
// 	}
// 	return []int{larger, smaller}
// }
