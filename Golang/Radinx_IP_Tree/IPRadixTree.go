/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/8/22 16:46 
  */

package Radinx_IP_Tree

import (
	"IPTree/Radinx_IP_Tree/utils"
	"errors"
	"sort"
)

// Nginx radix tree source code
// https://trac.nginx.org/nginx/browser/nginx/src/core/ngx_radix_tree.h
// https://trac.nginx.org/nginx/browser/nginx/src/core/ngx_radix_tree.c

// radix 节点结构体
type radix_node struct {
	left	*radix_node
	right	*radix_node
	parent	*radix_node
	value 	utils.Uint32Slice
}

// radix tree 结构体  对外开放
// free 和 alloc 配合，用于新节点内存申请
type Radix_tree struct {
	root 	*radix_node
	free 	*radix_node		// free 虽然指向一个节点，但是它的移动过程相当于空闲节点的链表
	alloc 	[]radix_node
}

// 创建一个新节点
func (tree *Radix_tree) newNode() (p *radix_node) {
	if tree.free != nil {
		p = tree.free
		tree.free = tree.free.right
		p.left = nil
		p.right = nil
		p.parent = nil
		p.value = nil
		return p
	}

	length := len(tree.alloc)
	if length == cap(tree.alloc) {
		tree.alloc = make([]radix_node, length+200)[:1]
		length = 0
	} else {
		tree.alloc = tree.alloc[:length+1]
	}
	return &(tree.alloc[length])
}

// 建立一个 Radix Tree
func Create_Radix_Tree(preallocate int) *Radix_tree {
	tree := new(Radix_tree)
	tree.root = tree.newNode()
	if preallocate == 0 {
		return tree
	}

	// 默认设置 preallocate 为6，在 nginx 中根据硬件平台做了判断
	if preallocate > 6 || preallocate < 0 {
		preallocate = 6
	}

	var mask uint32 = 0
	var ip  uint32 = 0
	var inc  uint32 = 0x80000000

	for ; preallocate > 0; preallocate -- {
		ip = 0
		mask >>= 1
		mask |= 0x80000000

		for ; ip != 0; {
			tree.radix_tree_insert(ip, mask, nil)
			ip += inc
		}
		inc >>= 1
	}
	return tree
}

func (tree *Radix_tree) radix_tree_insert(ip uint32, mask uint32, value utils.Uint32Slice) error {
	var bit uint32 = 0x80000000
	node := tree.root
	next := tree.root

	for ; bit & mask != 0;  {
		if ip & bit != 0 {
			next = node.right
		} else {
			next = node.left
		}
		if next == nil {
			break
		}
		bit >>= 1
		node = next
	}

	if next != nil {
		// 向 node.value 插入值
		node.value = append(node.value, value[:]...)
		// 去重+排序  先排序再去重
		sort.Sort(node.value)
		node.value = utils.UniqueUint32Slice(node.value)
		return nil
	}

	for ; bit & mask != 0;  {
		next = tree.newNode()
		if next == nil {
			return errors.New("Error: Create Node Failed")
		}
		next.left = nil
		next.right = nil
		next.parent = node
		next.value = nil

		if ip & bit != 0 {
			node.right = next
		} else {
			node.left = next
		}
		bit >>= 1
		node = next
	}
	// 向 node.value 插入值
	node.value = append(node.value, value[:]...)
	// 去重+排序  先排序再去重
	sort.Sort(node.value)
	node.value = utils.UniqueUint32Slice(node.value)
	return nil
}

func (tree *Radix_tree) Radix_Tree_Insert(cidr string, value utils.Uint32Slice) error {
	ip, mask, err := utils.ParseCidr4(cidr)
	if err != nil {
		return err
	}
	return tree.radix_tree_insert(ip, mask, value)
}

func (tree *Radix_tree) Radix_Tree_Delete(cidr string) error {
	ip, mask, err := utils.ParseCidr4(cidr)
	if err != nil {
		return err
	}

	var bit uint32 = 0x80000000
	node := tree.root
	for ; (node != nil) && (bit & mask != 0);  {
		if ip & bit != 0 {
			node = node.right
		} else {
			node = node.left
		}
		bit >>= 1
	}
	if node == nil {
		return errors.New("Error: Wrong Mask")
	}
	if node.right != nil || node.left != nil {
		if node.value != nil {
			node.value = nil
			return nil
		}
		return errors.New("Error: Wrong Cidr")
	}
	for {
		if node.parent.right == node {
			node.parent.right = nil
		} else {
			node.parent.left = nil
		}
		node.right = tree.free
		tree.free = node

		node = node.parent
		if node.right != nil || node.left != nil {
			break
		}
		if node.value != nil {
			break
		}
		if node.parent == nil {
			break
		}
	}
	return nil
}

// 查找一个ip，这里的输入应为 192.168.1.1/32
func (tree *Radix_tree) Radix_Tree_Search(cidr string) []uint32 {
	ip, _, err := utils.ParseCidr4(cidr)
	if err != nil {
		return nil
	}
	var bit uint32 = 0x80000000
	node := tree.root
	var value utils.Uint32Slice

	for ; node != nil;  {
		if node.value != nil {
			// 把节点的 value 装入 value
			value = append(value, node.value[:]...)
		}
		if ip & bit != 0 {
			node = node.right
		} else {
			node = node.left
		}
		bit >>= 1
	}
	sort.Sort(value)
	return value
}