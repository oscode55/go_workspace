/*
* @Author: myname
* @Date:   2018-10-19 14:45:33
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:24:01
*/
package conversions

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
//类型转换 将double转换成其他类型
// Convert double to float
type D2F struct{ base.NoOperandsInstruction }

func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

// Convert double to int
type D2I struct{ base.NoOperandsInstruction }

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// Convert double to long
type D2L struct{ base.NoOperandsInstruction }

func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
