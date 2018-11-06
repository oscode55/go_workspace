/*
* @Author: myname
* @Date:   2018-10-19 15:01:00
* @Last Modified by:   myname
* @Last Modified time: 2018-10-19 15:02:26
 */
package base

import "jvmgo/ch05/rtda"

//跳转 传入pc偏移量
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
