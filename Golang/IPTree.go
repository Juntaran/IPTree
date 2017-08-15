/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/8/14 11:40 
  */

package IPTree

import "fmt"

// 路由元素节点
type routeNode struct {
	left	*routeNode
	right	*routeNode
	port	int
}

// 路由树结构
type routeTree struct {
	root	*routeNode
	Size 	uint
}

// 创建节点
func CreateNode() *routeNode {
	pNode := new(routeNode)
	pNode.left = nil
	pNode.right = nil
	pNode.port = -1
	return pNode
}

// 创建一颗路由树
func CreateRouteTree(node *routeNode) *routeTree {
	if node == nil {
		return &routeTree{}
	}
	return &routeTree{
		root: 	node,
		Size:	1,
	}
}

// 创建路由表
// 在这个路由表中，0向左，1向右
// 输入：路由IP 掩码位 端口号
func (tree *routeTree) InsertRoute(iRoute int, iMask uint, iPort int) {
	var judge = 0
	node := CreateNode()
	if tree.Size == 0 {
		tree.root = node
		tree.Size ++
		return
	}
	root := tree.root
	var currentNode *routeNode
	currentNode = root
	var i uint
	for i = 0; i < iMask; i++ {
		// 根据ip二进制从左到右按位解析
		judge = (iRoute >> (31-i)) & 0x1

		if judge == 0 {
			if currentNode.left == nil {
				currentNode.left = CreateNode()
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				currentNode.right = CreateNode()
			}
			currentNode = currentNode.right
		}
	}
	tree.Size ++
	currentNode.port = iPort
}

// 定位路由节点
func (tree *routeTree) LocateRoute(ip int) *routeNode {
	var judge int = 0
	root := tree.root
	var currentNode *routeNode = root
	if currentNode.port != -1 {
	}
	fmt.Printf("start locate ip: %32b\n", ip)
	var i uint
	for i = 0; i < 32; i++ {
		// 根据ip二进制从左到右按位解析
		judge = (ip >> (31-i)) & 0x1
		if judge == 0 {
			if currentNode.left != nil {
				currentNode = currentNode.left
			} else {
				break
			}
		} else {
			if currentNode.right != nil {
				currentNode = currentNode.right
			} else {
				break
			}
		}
	}
	fmt.Println(currentNode.port)
	return currentNode
}

// 删除一条路由
func (tree *routeTree) DeleteRoute(iRoute int, iMask uint) {
	judge := 0
	fmt.Printf("Delete route: %32b, mask: %4d\n", iRoute, iMask)
	root := tree.root
	var currentNode *routeNode = root
	var i uint
	for i = 0; i < iMask; i++ {
		// 根据ip二进制从左到右按位解析
		judge = (iRoute >> (31-i)) & 0x1
		if judge == 0 {
			if currentNode.left == nil {
				currentNode.left = CreateNode()
			}
			currentNode = currentNode.left
		} else {
			if currentNode.right == nil {
				currentNode.right = CreateNode()
			}
			currentNode = currentNode.right
		}
	}
	if i < iMask {
		fmt.Println("Delete Error.")
		return
	}
	currentNode.port = -1
	currentNode.left = nil
	currentNode.right = nil
	tree.Size --
	fmt.Println("Delete Success")
	return
}

// 查找路由表函数
func (tree *routeTree)SearchRoute(ip int) int {
	currentNode := tree.LocateRoute(ip)
	if currentNode.port == -1 {
		fmt.Printf("Not in rule, %d\n", currentNode.port)
	}   else {
		fmt.Printf("Get %d\n", currentNode.port)
	}
	return currentNode.port
}

// 层序遍历
func (tree *routeTree)LevelTraverse()  {
	fmt.Println("Start LevelTraverse")
	if tree.root == nil {
		return
	}
	root := tree.root
	var treeQueue []*routeNode
	treeQueue = append(treeQueue, root)
	var current int = 0
	var last int = 1
	for current < len(treeQueue) {
		last = len(treeQueue)
		for current < last {
			fmt.Printf("%4d ", treeQueue[current].port)
			if treeQueue[current].left != nil {
				treeQueue = append(treeQueue, treeQueue[current].left)
			}
			if treeQueue[current].right != nil {
				treeQueue = append(treeQueue, treeQueue[current].right)
			}
			current ++
		}
		fmt.Println()
	}
	fmt.Println()
}

// 显示当前routeTree中的规则数量
func (tree *routeTree)RouteRuleNumber()  {
	fmt.Println("The Rule Number is", tree.Size - 1)
}