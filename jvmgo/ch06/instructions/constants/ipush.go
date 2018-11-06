/*
* @Author: myname
* @Date:   2018-10-18 15:15:55
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:23:30
 */
package constants

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

type BIPUSH struct { //Push byte
	//从栈中获取一个byte类型 拓展成int 入栈
	val int8
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

type SIPUSH struct { //Push short
	//从栈中获取一个short类型 拓展成int 入栈
	val int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
