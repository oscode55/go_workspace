/*
* @Author: myname
* @Date:   2018-10-19 16:13:36
* @Last Modified by:   myname
* @Last Modified time: 2018-10-19 16:15:12
 */
package extended

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

//根据引用是否是null进行跳转
// Branch if reference is null
type IFNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

// Branch if reference not null
type IFNONNULL struct{ base.BranchInstruction }

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
