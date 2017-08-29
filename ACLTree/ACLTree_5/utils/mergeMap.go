/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/8/29 18:15 
  */

package utils

import "sort"

// 合并两个map
// 1. 用一个 slice 存储 map1 和 map2 重复的 key
// 2. 合并两个 map
// 3. 遍历合并后的 map 对 value 去重+排序
func MergeMap(map1, map2 map[uint32]Uint32Slice) map[uint32]Uint32Slice {
	var key1, key2 Uint32Slice
	for k1, _ := range map1 {
		key1 = append(key1, k1)
	}
	for k2, _ := range map2 {
		key2 = append(key2, k2)
	}
	if len(key1) == 0 {
		return map2
	}
	if len(key2) == 0 {
		return map1
	}

	map3 := make(map[uint32]Uint32Slice, 10)

	for i := 0; i < len(key1); i++ {
		for j := 0; j < len(key2); j++ {
			// 把相同的 key/value 插入到 key3 中
			if key1[i] == key2[j] {
				map3[key1[i]] = map1[key1[i]]
			}
		}
	}

	// 把所有的 map2 放入到 map1 中
	// 如果 key 相同， map2 会覆盖 map1
	for i := 0; i < len(key2); i++ {
		map1[key2[i]] = map2[key2[i]]
		// 去重 + 排序
		sort.Sort(map1[key2[i]])
		map1[key2[i]] = UniqueUint32Slice(map1[key2[i]])
	}

	// 把 map3 中的元素再放入 map1 中
	if len(map3) > 0 {
		for k, v := range map3 {
			map1[k] = append(map1[k], v[:]...)
			// 去重 + 排序
			sort.Sort(map1[k])
			map1[k] = UniqueUint32Slice(map1[k])
		}
	}
	return map1
}
