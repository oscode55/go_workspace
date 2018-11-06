/*
* @Author: myname
* @Date:   2018-10-19 15:14:06
* @Last Modified by:   myname
* @Last Modified time: 2018-10-19 15:14:46
*/
package control

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"
/**
 *无条件跳转
 */
// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
