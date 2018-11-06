/*
* @Author: myname
* @Date:   2018-10-19 15:08:08
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:27:39
*/
package comparisons

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
//弹出栈顶两个引用 根据是否相同进行跳转
// Branch if reference comparison succeeds
type IF_ACMPEQ struct{ 
	base.BranchInstruction 
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}
