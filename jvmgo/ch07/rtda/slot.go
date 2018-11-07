/*
* @Author: myname
* @Date:   2018-10-18 10:19:35
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:31:00
 */
package rtda

import "jvmgo/ch07/rtda/heap"

//double 和 long 型的占据两个Slot
type Slot struct {
	// boolean/type/char/short/int/float/reference/returnAddress
	num int32        //存储32位及以内的
	ref *heap.Object //引用类型
}
