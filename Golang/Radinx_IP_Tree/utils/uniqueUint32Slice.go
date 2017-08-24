/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/8/24 18:44 
  */

package utils

func UniqueUint32Slice(slice Uint32Slice) Uint32Slice {
	length := len(slice)

	if length <= 1 {
		return slice
	}

	var ret Uint32Slice
	for i := 0; i < length-1;  {
		ret = append(ret, slice[i])
		for j := i+1; j < length; j++ {
			if slice[i] == slice[j] {
				continue
			} else {
				i = j
				break
			}
		}
	}
	if slice[length-2] == slice[length-1] {
		return ret
	}
	ret = append(ret, slice[length-1])
	return ret
}

