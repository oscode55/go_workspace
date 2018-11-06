/*
* @Author: myname
* @Date:   2018-10-18 09:53:53
* @Last Modified by:   myname
* @Last Modified time: 2018-10-19 16:33:09
 */
package rtda

import "jvmgo/ch06/rtda/heap"

type Thread struct {
	pc    int
	stack *Stack //jvm Stack Pointer
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024), //可以容纳1024帧
	}
}
func (self *Thread) PC() int { //getter
	return self.pc
}
func (self *Thread) SetPC(pc int) { //setter
	self.pc = pc
}
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

//返回当前帧 即栈顶
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}
