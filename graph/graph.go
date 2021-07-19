package graph

import "sync"

type Node struct {
	Value int
}

// Test
// g := graph.NewGraph()
// n1, n2, n3, n4, n5 := graph.Node{1}, graph.Node{2}, graph.Node{3}, graph.Node{4}, graph.Node{5}

// g.AddNode(&n1)
// g.AddNode(&n2)
// g.AddNode(&n3)
// g.AddNode(&n4)
// g.AddNode(&n5)

// g.AddEdge(&n1, &n2)
// g.AddEdge(&n1, &n5)
// g.AddEdge(&n2, &n3)
// g.AddEdge(&n2, &n4)
// g.AddEdge(&n2, &n5)
// g.AddEdge(&n3, &n4)
// g.AddEdge(&n4, &n5)
// fmt.Println(g.Edges)
// g.BFS(&n4, func(node *graph.Node) {
// 	fmt.Println(node)
// })
type Graph struct {
	Nodes []*Node           //记录节点
	Edges map[*Node][]*Node //记录边
	Lock  sync.RWMutex      //读写锁
}

func NewGraph() *Graph {
	return &Graph{
		Edges: make(map[*Node][]*Node),
	}
}

func (g *Graph) AddNode(node *Node) {
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) AddEdge(node1, node2 *Node) {
	g.Edges[node1] = append(g.Edges[node1], node2)
}

//广度遍历，参数是对节点的操作
func (g *Graph) BFS(start *Node, f func(node *Node)) {
	g.Lock.RLock()
	defer g.Lock.RUnlock()
	//标记节点是否已经被访问过
	visited := make(map[*Node]bool)
	//新建队列，并把起点压入队列
	var nodeList []*Node
	nodeList = append(nodeList, start)
	for {
		if len(nodeList) == 0 { //遍历完毕，退出
			break
		}
		//取出头部node
		visitNode := nodeList[0]
		nodeList = nodeList[1:]
		//标记
		visited[visitNode] = true
		for _, node := range g.Edges[visitNode] { //把连接的节点压入队列
			if visited[node] { //如果已经被遍历就继续
				continue
			}
			//压入队列
			nodeList = append(nodeList, node)
			//标记
			visited[node] = true
		}
		//对节点调用回调函数
		f(visitNode)
	}
}

// //迷路的机器人，深搜找终点
// //优化点：状态压缩，二维切片优化成普通切片(行列最大不超过100，可以以101为区间大小，给切片分段)
// func pathWithObstacles(obstacleGrid [][]int) [][]int {
// 	var (
// 		result [][]int
// 		rowEnd = len(obstacleGrid) - 1
// 		colEnd = len(obstacleGrid[0]) - 1
// 		dfs    func([][]int) //内部函数递归需要用到该变量
// 	)
// 	dfs = func(path [][]int) {
// 		if len(result) != 0 { //已找到路径
// 			return
// 		}
// 		row, col := path[len(path)-1][0], path[len(path)-1][1]
// 		if obstacleGrid[row][col] == 1 { //判断路径是否可通
// 			return
// 		}
// 		if row == rowEnd && col == colEnd {
// 			result = make([][]int, rowEnd+colEnd+1) //结果的长度必定是长加高-1
// 			copy(result, path)
// 		}
// 		obstacleGrid[row][col] = 1
// 		if row < rowEnd { //向下走
// 			dfs(append(path, []int{row + 1, col})) //采用append的方式可以不改变path的值，让path仍然记录当前位置
// 		}
// 		if col < colEnd { //向右走
// 			dfs(append(path, []int{row, col + 1}))
// 		}
// 	}
// 	dfs([][]int{{0, 0}})
// 	return result
// }

//迷路的机器人，深搜找终点
func pathWithObstacles(obstacleGrid [][]int) [][]int {
	var (
		result [][]int
		rowEnd = len(obstacleGrid) - 1
		colEnd = len(obstacleGrid[0]) - 1
		dfs    func([]int) //内部函数递归需要用到该变量
		m      int         = 101
	)
	dfs = func(path []int) {
		if len(result) != 0 { //已找到路径
			return
		}
		now := path[len(path)-1]
		row, col := now/m, now%m
		if obstacleGrid[row][col] == 1 { //判断路径是否可通
			return
		}
		if row == rowEnd && col == colEnd {
			result = make([][]int, rowEnd+colEnd+1) //结果的长度必定是长加高-1
			for t, p := range path {
				result[t] = []int{p / m, p % m}
			}
		}
		obstacleGrid[row][col] = 1
		if row < rowEnd { //向下走
			dfs(append(path, now+m)) //采用append的方式可以不改变path的值，让path仍然记录当前位置
		}
		if col < colEnd { //向右走
			dfs(append(path, now+1))
		}
	}
	dfs([]int{0})
	return result
}

// 颜色填充
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var (
		flag     map[int]bool = make(map[int]bool) //标记节点时候已经被染色
		dfs      func(tempSr, tempSc int)
		rLen     = len(image)
		cLen     = len(image[0])
		oldColor = image[sr][sc] //记录原始节点颜色
	)
	dfs = func(tempSr, tempSc int) { //深度搜索
		if image[tempSr][tempSc] != oldColor {
			return
		}
		image[tempSr][tempSc] = newColor
		flag[tempSr*cLen+tempSc] = true
		if 0 <= tempSr-1 && !flag[(tempSr-1)*cLen+tempSc] { //上
			dfs(tempSr-1, tempSc)
		}
		if tempSr+1 < rLen && !flag[(tempSr+1)*cLen+tempSc] { //下
			dfs(tempSr+1, tempSc)
		}
		if 0 <= tempSc-1 && !flag[tempSr*cLen+(tempSc-1)] { //左
			dfs(tempSr, tempSc-1)
		}
		if tempSc+1 < cLen && !flag[tempSr*cLen+(tempSc+1)] { //右
			dfs(tempSr, tempSc+1)
		}
	}
	dfs(sr, sc)
	return image
}
