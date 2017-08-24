/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2017/8/24 17:17 
  */

package utils

// Go Sort interface
//定义interface{},并实现sort.Interface接口的三个方法
type Uint32Slice []uint32

func (c Uint32Slice) Len() int {
	return len(c)
}
func (c Uint32Slice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c Uint32Slice) Less(i, j int) bool {
	return c[i] < c[j]
}


//func main() {
//	a := Uint32Slice{1, 3, 5, 7, 2}
//	b := []float64{1.1, 2.3, 5.3, 3.4}
//	c := []int{1, 3, 5, 4, 2}
//	fmt.Println(sort.IsSorted(a)) //false
//	if !sort.IsSorted(a) {
//		sort.Sort(a)
//	}
//
//	if !sort.Float64sAreSorted(b) {
//		sort.Float64s(b)
//	}
//	if !sort.IntsAreSorted(c) {
//		sort.Ints(c)
//	}
//	fmt.Println(a)//[1 2 3 5 7]
//	fmt.Println(b)//[1.1 2.3 3.4 5.3]
//	fmt.Println(c)// [1 2 3 4 5]
//}
