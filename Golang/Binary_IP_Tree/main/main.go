/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/8/14 13:47
 */

package main

import (
	IPTree "IPTree/Binary_IP_Tree"
	"fmt"
)

func main() {
	a := make([]int, 4)
	var Route int = 0
	var Mask uint = 0
	var Port int = 0
	var n int

	fmt.Scanf("%d", &n)
	fmt.Println(n)

	rootNode := IPTree.CreateNode()
	routeTree := IPTree.CreateRouteTree(rootNode)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d.%d.%d.%d/%d %d\n", &a[0], &a[1], &a[2], &a[3], &Mask, &Port)
		fmt.Printf("%d.%d.%d.%d/%d %d\n", a[0], a[1], a[2], a[3], Mask, Port)
		Route = a[0]<<24 | a[1]<<16 | a[2]<<8 | a[3]<<0
		fmt.Printf("%32b\n", Route)
		routeTree.InsertRoute(Route, Mask, Port)
	}
	routeTree.RouteRuleNumber()
	routeTree.LevelTraverse()

	fmt.Println("Input Ip you want Search")
	ip := make([]int, 4)
	fmt.Scanf("%d.%d.%d.%d\n", &ip[0], &ip[1], &ip[2], &ip[3])
	iIp := ip[0]<<24 | ip[1]<<16 | ip[2]<<8 | ip[3]<<0
	routeTree.SearchRoute(iIp)

	// 删除ip这条路由 掩码为/32
	routeTree.DeleteRoute(iIp, 32)
	routeTree.LevelTraverse()
	routeTree.SearchRoute(iIp)
	routeTree.RouteRuleNumber()
}
