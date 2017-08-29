/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/8/25 11:00
  */

package main

import (
	"IPTree/ACLTree"
	"IPTree/ACLTree/utils"
	"log"
	"fmt"
)

func main() {
	aclTree := ACLTree.Create_ACL_Tree(0)
	err := aclTree.ACL_Tree_Insert("192.168.1.0/24", []uint32{1,2,3,4,5}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	white, black := aclTree.ACL_Tree_Search("192.168.1.111")
	fmt.Println(white, black)

	value1 := utils.Uint32Slice{111,121}
	err = aclTree.ACL_Tree_Insert("192.168.0.0/16", value1, nil)
	if err != nil {
		log.Fatalln(err)
	}
	white, black = aclTree.ACL_Tree_Search("192.168.1.111")
	fmt.Println(white, black)

	value2 := utils.Uint32Slice{3,7,6,5}
	err = aclTree.ACL_Tree_Insert("192.168.1.0/24", value2, []uint32{8})
	if err != nil {
		log.Fatalln(err)
	}
	white, black = aclTree.ACL_Tree_Search("192.168.1.111")
	fmt.Println(white, black)

	err = aclTree.ACL_Tree_Delete("192.168.1.0/24")
	if err != nil {
		log.Fatalln(err)
	}

	white, black = aclTree.ACL_Tree_Search("192.168.1.111")
	fmt.Println(white, black)

	err = aclTree.ACL_Tree_Insert_Lot("192.168.0.0/16", "100-120", "199-203")
	if err != nil {
		log.Fatalln(err)
	}
	white, black = aclTree.ACL_Tree_Search("192.168.1.111")
	fmt.Println(white, black)
}
