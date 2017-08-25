/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/8/25 10:12
  */

package ACLTree

import (
	"errors"
	"sort"
	"strconv"
	"IPTree/ACLTree/utils"
)

// Nginx radix tree source code
// https://trac.nginx.org/nginx/browser/nginx/src/core/ngx_radix_tree.h
// https://trac.nginx.org/nginx/browser/nginx/src/core/ngx_radix_tree.c

// acl 节点结构体
type acl_node struct {
	left		*acl_node
	right		*acl_node
	parent		*acl_node
	white 		utils.Uint32Slice		// 白名单 port_list
	black		utils.Uint32Slice		// 黑名单 port_list
}

// acl tree 结构体  对外开放
// free 和 alloc 配合，用于新节点内存申请
type ACL_tree struct {
	root 	*acl_node
	free 	*acl_node		// free 虽然指向一个节点，但是它的移动过程相当于空闲节点的链表
	alloc 	[]acl_node
}

// 创建一个新节点
func (tree *ACL_tree) newNode() (p *acl_node) {
	if tree.free != nil {
		p = tree.free
		tree.free = tree.free.right
		p.left = nil
		p.right = nil
		p.parent = nil
		p.white = nil
		p.black = nil
		return p
	}

	length := len(tree.alloc)
	if length == cap(tree.alloc) {
		tree.alloc = make([]acl_node, length+200)[:1]
		length = 0
	} else {
		tree.alloc = tree.alloc[:length+1]
	}
	return &(tree.alloc[length])
}

// 建立一个 ACL Tree
func Create_ACL_Tree(preallocate int) *ACL_tree {
	tree := new(ACL_tree)
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
			tree.acl_tree_insert(ip, mask, nil, nil)
			ip += inc
		}
		inc >>= 1
	}
	return tree
}

func (tree *ACL_tree) acl_tree_insert(ip uint32, mask uint32, white utils.Uint32Slice, black utils.Uint32Slice) error {
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
		var tempwhite utils.Uint32Slice = node.white
		var tempblack utils.Uint32Slice = node.black

		tempwhite = append(tempwhite, white[:]...)
		tempblack = append(tempblack, black[:]...)

		sort.Sort(tempwhite)
		sort.Sort(tempblack)

		tempwhite = utils.UniqueUint32Slice(tempwhite)
		tempblack = utils.UniqueUint32Slice(tempblack)

		// 插入后校验黑白名单是否冲突
		if len(tempwhite) != 0 && len(tempblack) != 0 {
			for i := 0; i <  len(tempwhite); i++ {
				for j := 0; j < len(tempblack); j++ {
					if tempwhite[i] == tempblack[j] {
						// 如果冲突了，直接 error ，不会对节点值有影响
						return errors.New("Error: White and Black Conflict")
					}
				}
			}
		}
		node.white = tempwhite
		node.black = tempblack

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
		next.white = nil
		next.black = nil

		if ip & bit != 0 {
			node.right = next
		} else {
			node.left = next
		}
		bit >>= 1
		node = next
	}
	// 向 node.value 插入值
	var tempwhite utils.Uint32Slice = node.white
	var tempblack utils.Uint32Slice = node.black

	tempwhite = append(tempwhite, white[:]...)
	tempblack = append(tempblack, black[:]...)

	// 去重+排序  先排序再去重
	sort.Sort(tempwhite)
	sort.Sort(tempblack)

	tempwhite = utils.UniqueUint32Slice(tempwhite)
	tempblack = utils.UniqueUint32Slice(tempblack)

	// 插入后校验黑白名单是否冲突
	if len(tempwhite) != 0 && len(tempblack) != 0 {
		for i := 0; i <  len(tempwhite); i++ {
			for j := 0; j < len(tempblack); j++ {
				if tempwhite[i] == tempblack[j] {
					// 如果冲突了，直接 error ，不会对节点值有影响
					return errors.New("Error: White and Black Conflict")
				}
			}
		}
	}
	node.white = tempwhite
	node.black = tempblack
	return nil
}

func (tree *ACL_tree) ACL_Tree_Insert(cidr string, white utils.Uint32Slice, black utils.Uint32Slice) error {
	// 插入的时候，黑白名单不能冲突
	if len(white) != 0 && len(black) != 0 {
		for i := 0; i <  len(white); i++ {
			for j := 0; j < len(black); j++ {
				if white[i] == black[j] {
					return errors.New("Error: White and Black Conflict")
				}
			}
		}
	}
	ip, mask, err := utils.ParseCidr4(cidr)
	if err != nil {
		return err
	}
	return tree.acl_tree_insert(ip, mask, white, black)
}

// 插入的黑白名单为一段 例如 tree.ACL_Tree_Insert_Lot(192.168.1.0/24, 1-100, 101-200)
func (tree *ACL_tree) ACL_Tree_Insert_Lot(cidr string, white string, black string) error {
	// 检查插入的黑白名单格式是否正确
	// 插入的黑白名单只能有两种形式:
	// 1. 单个数字 100
	// 2. 端口范围 100-200
	var num1, num2 int
	var pos1, pos2 int
	for i := 0; i < len(white); {
		if white[i] >= '0' && white[i] <= '9' {
			i ++
		} else if white[i] == '-' {
			if i == 0 || i == len(white)-1 {
				return errors.New("Error: Wrong Input")
			}
			num1 ++
			pos1 = i
			if num1 > 1 {
				return errors.New("Error: Wrong Input")
			}
			i ++
		} else {
			return errors.New("Error: Wrong Input")
		}
	}
	for i := 0; i < len(black); {
		if black[i] >= '0' && black[i] <= '9' {
			i ++
		} else if black[i] == '-' {
			if i == 0 || i == len(black)-1 {
				return errors.New("Error: Wrong Input")
			}
			num2 ++
			pos2 = i
			if num2 > 1 {
				return errors.New("Error: Wrong Input")
			}
			i ++
		} else {
			return errors.New("Error: Wrong Input")
		}
	}
	var whiteStart, whiteEnd, blackStart, blackEnd int
	var whiteTemp []uint32
	var blackTemp []uint32
	if num1 == 0 && num2 == 0 {
		// 输入的黑白名单都是单独的数字
		whiteTemp = utils.StringToUint32Slice(white)
		blackTemp = utils.StringToUint32Slice(black)
	}
	if num1 == 1 && num2 == 0 {
		// 白名单为端口段，黑名单为单独的数字
		whiteStart, _ = strconv.Atoi(white[:pos1])
		whiteEnd, _ = strconv.Atoi(white[pos1+1:])
		for i := whiteStart; i <= whiteEnd; i++ {
			whiteTemp = append(whiteTemp, uint32(i))
		}
		blackTemp = utils.StringToUint32Slice(black)
	}
	if num1 == 0 && num2 == 1 {
		// 白名单为单独的数字，黑名单为端口段
		blackStart, _ = strconv.Atoi(black[:pos2])
		blackEnd, _ = strconv.Atoi(black[pos2+1:])
		for i := blackStart; i <= blackEnd; i++ {
			blackTemp = append(blackTemp, uint32(i))
		}
		whiteTemp = utils.StringToUint32Slice(white)
	}
	if num1 == 1 && num2 == 1 {
		// 黑白名单均为端口段
		whiteStart, _ = strconv.Atoi(white[:pos1])
		whiteEnd, _ = strconv.Atoi(white[pos1+1:])
		blackStart, _ = strconv.Atoi(black[:pos2])
		blackEnd, _ = strconv.Atoi(black[pos2+1:])

		for i := whiteStart; i <= whiteEnd; i++ {
			whiteTemp = append(whiteTemp, uint32(i))
		}
		for i := blackStart; i <= blackEnd; i++ {
			blackTemp = append(blackTemp, uint32(i))
		}
	}
	return tree.ACL_Tree_Insert(cidr, whiteTemp, blackTemp)
}

func (tree *ACL_tree) ACL_Tree_Delete(cidr string) error {
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
		if node.white != nil || node.black != nil {
			node.white = nil
			node.black = nil
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
		if node.white != nil || node.black != nil {
			break
		}
		if node.parent == nil {
			break
		}
	}
	return nil
}

// 查找一个ip，这里的输入应为 192.168.1.1/32
func (tree *ACL_tree) ACL_Tree_Search(sip string) ([]uint32, []uint32) {
	cidr := sip + "/32"
	ip, _, err := utils.ParseCidr4(cidr)
	if err != nil {
		return nil, nil
	}
	var bit uint32 = 0x80000000
	node := tree.root
	var white utils.Uint32Slice
	var black utils.Uint32Slice

	for ; node != nil;  {
		if node.white != nil || node.black != nil {
			// 把节点的 value 装入 value
			white = append(white, node.white[:]...)
			black = append(black, node.black[:]...)
		}
		if ip & bit != 0 {
			node = node.right
		} else {
			node = node.left
		}
		bit >>= 1
	}
	sort.Sort(white)
	sort.Sort(black)
	return white, black
}