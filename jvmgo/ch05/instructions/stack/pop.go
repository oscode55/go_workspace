/*
* @Author: myname
* @Date:   2018-10-19 11:53:44
* @Last Modified by:   myname
* @Last Modified time: 2018-10-19 17:08:29
 */
package stack

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// pop 弹出一个Slot
type POP struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// pop2 弹出两个Slot
type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
