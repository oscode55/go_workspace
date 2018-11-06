/*
* @Author: myname
* @Date:   2018-10-18 10:19:35
* @Last Modified by:   myname
* @Last Modified time: 2018-10-18 14:00:06
*/
package rtda
//double 和 long 型的占据两个Slot
type Slot struct {
	// boolean/type/char/short/int/float/reference/returnAddress
	num int32 //存储32位及以内的 
	ref *Object //引用类型
}