/*
* @Author: myname
* @Date:   2018-10-18 10:09:24
* @Last Modified by:   myname
* @Last Modified time: 2018-10-18 11:39:22
 */
package rtda

type Frame struct {
	next         *Frame        //链表数据结构
	localVars    LocalVars     //局部变量表指针
	operandStack *OperandStack //操作数栈指针
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
func (self Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self Frame) OperandStack() *OperandStack {
	return self.operandStack
}
