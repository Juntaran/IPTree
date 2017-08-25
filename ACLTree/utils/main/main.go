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
	a := "123454"
	b := utils.StringToUint32Slice(a)
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}
