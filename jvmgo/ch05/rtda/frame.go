/*
* @Author: myname
* @Date:   2018-10-18 10:09:24
* @Last Modified by:   myname
* @Last Modified time: 2018-10-19 16:57:26
*/
package rtda
type Frame struct {
	next *Frame //链表数据结构
	localVars LocalVars //局部变量表指针
	operandStack *OperandStack //操作数栈指针
	thread *Thread 
	nextPC int
}
func NewFrame(thread *Thread,maxLocals,maxStack uint) *Frame {
	return &Frame {
		thread: thread,
		localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
//settersand getters
func (self Frame) LocalVars() LocalVars{
	return self.localVars
}
func (self Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}