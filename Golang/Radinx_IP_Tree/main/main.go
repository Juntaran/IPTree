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
	value = radixTree.Radix_Tree_Search("192.168.1.111/32")
	fmt.Println(value)
}
