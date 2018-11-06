/*
* @Author: myname
* @Date:   2018-10-19 14:52:40
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:27:40
 */
package comparisons

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

// Compare float
//浮点型 除了大于 小于 等于 还有第四种结果 无法比较
//因为浮点数可能产出NaN (not a Number)
type FCMPG struct{ base.NoOperandsInstruction }

func (self *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct{ base.NoOperandsInstruction }

func (self *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
