/*
* @Author: myname
* @Date:   2018-10-19 14:49:40
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:23:17
*/
package comparisons

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

// Compare long
type LCMP struct{ base.NoOperandsInstruction }

//栈顶两个long变量弹出 比较 结果(0,1,-1)入栈
func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
