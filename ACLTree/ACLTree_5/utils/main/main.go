/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/8/25 15:15 
  */

package main

import (
	"fmt"
	"IPTree/ACLTree/utils"
	"reflect"
)

func main() {

	// uint32slice test
	a := "123454"
	b := utils.StringToUint32Slice(a)
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	// uint32slice test end

	// map test
	map1 := map[uint32]utils.Uint32Slice{1:{3,2,2,1}, 2:{3,4,5,2,1,3}}
	map2 := map[uint32]utils.Uint32Slice{2:{1,2,3}, 3:{5,6,1,7}}
	map1 = utils.MergeMap(map1, map2)
	fmt.Println(map1)
	// map test end


	fmt.Println(map1 == nil)
	map1 = nil
	fmt.Println(map1)
}
