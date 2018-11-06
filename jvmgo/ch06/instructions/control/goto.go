/*
* @Author: myname
* @Date:   2018-10-19 15:14:06
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:23:42
 */
package control

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

/**
 *无条件跳转
 */
// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
