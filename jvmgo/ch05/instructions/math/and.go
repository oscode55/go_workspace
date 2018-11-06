/*
* @Author: myname
* @Date:   2018-10-19 14:40:07
* @Last Modified by:   myname
* @Last Modified time: 2018-10-19 14:46:49
*/
package math

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Boolean AND int
type IAND struct{ base.NoOperandsInstruction }

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

// Boolean AND long
type LAND struct{ base.NoOperandsInstruction }

func (self *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
