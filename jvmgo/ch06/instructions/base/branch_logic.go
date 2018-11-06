/*
* @Author: myname
* @Date:   2018-10-19 15:01:00
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:22:20
*/
package base

import "jvmgo/ch06/rtda"

//跳转 传入pc偏移量
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
