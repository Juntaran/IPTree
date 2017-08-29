/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/8/25 15:10 
  */

package utils

import "strconv"

func StringToUint32Slice(test string) []uint32 {
	var ret []uint32
	k, _ := strconv.Atoi(test)
	var j uint32 = uint32(k)
	ret = append(ret, j)
	return ret
}