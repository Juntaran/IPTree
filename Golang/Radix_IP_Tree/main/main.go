/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/8/24 21:03 
  */

package main

import (
	"IPTree/Radinx_IP_Tree"
	"IPTree/Radinx_IP_Tree/utils"
	"log"
	"fmt"
)

func main() {
	radixTree := Radinx_IP_Tree.Create_Radix_Tree(0)
	value := utils.Uint32Slice{1,2,3,4}
	err := radixTree.Radix_Tree_Insert("192.168.1.0/24", value)
	if err != nil {
		log.Fatalln(err)
	}
	value3 := radixTree.Radix_Tree_Search("192.168.1.111/32")
	fmt.Println(value3)

	value1 := utils.Uint32Slice{111,121}
	err = radixTree.Radix_Tree_Insert("192.168.0.0/16", value1)
	if err != nil {
		log.Fatalln(err)
	}
	value3 = radixTree.Radix_Tree_Search("192.168.1.111/32")
	fmt.Println(value3)

	value2 := utils.Uint32Slice{3,7,6,5}
	err = radixTree.Radix_Tree_Insert("192.168.1.0/24", value2)
	if err != nil {
		log.Fatalln(err)
	}
	value3 = radixTree.Radix_Tree_Search("192.168.1.111/32")
	fmt.Println(value3)

	err = radixTree.Radix_Tree_Delete("192.168.1.0/24")
	if err != nil {
		log.Fatalln(err)
	}

	value3 = radixTree.Radix_Tree_Search("192.168.1.111/32")
	fmt.Println(value3)
}
