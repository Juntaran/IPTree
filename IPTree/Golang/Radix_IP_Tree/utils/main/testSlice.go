/**
 * Author: Juntaran
 * Email:  Jacinthmail@gmail.com
 * Date:   2017/8/24 17:04
 */

package main

import (
	"IPTree/IPTree/Golang/Radix_IP_Tree/utils"
	"fmt"
	"sort"
)

func main() {
	var a utils.Uint32Slice

	b := utils.Uint32Slice{4,12,3}
	fmt.Println(b)

	var i uint32
	fmt.Println(i)
	for i = 5; i > 0; i-- {
		a = append(a, i)
	}
	a = append(a, 0)

	fmt.Println(a)
	sort.Sort(a)
	fmt.Println(a)

	a = append(a[:2], a[3:]...)
	fmt.Println(a)

	b = append(b, a[:]...)
	fmt.Println(b)

	// 排序
	sort.Sort(b)
	fmt.Println(b)

	// 去重
	b = utils.UniqueUint32Slice(b)
	fmt.Println(b)
}
