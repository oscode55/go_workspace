/*
* @Author: myname
* @Date:   2018-10-18 15:06:44
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:23:36
 */
package constants

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

//nop指令 代表什么也不做
type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}
