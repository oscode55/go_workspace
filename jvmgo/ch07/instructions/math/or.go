/*
* @Author: myname
* @Date:   2018-10-19 16:51:04
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:27:36
 */
package math

import "jvmgo/ch07/instructions/base"
import "jvmgo/ch07/rtda"

// Boolean OR int
type IOR struct{ base.NoOperandsInstruction }

func (self *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

// Boolean OR long
type LOR struct{ base.NoOperandsInstruction }

func (self *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
