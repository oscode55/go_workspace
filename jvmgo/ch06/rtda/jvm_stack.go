/*
* @Author: myname
* @Date:   2018-10-18 10:00:30
* @Last Modified by:   myname
* @Last Modified time: 2018-10-18 13:53:31
*/
package rtda
type Stack struct {
	maxSize uint //栈容量(帧)
	size uint //当前栈大小(帧)
	_top *Frame //链表结构
}
func newStack(capacity uint) *Stack {
	return &Stack {
		maxSize: capacity,
	}
}
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.next = self._top
	}
	self._top = frame 
	self.size++
}
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = top.next
	top.next = nil
	self.size--
	return top 
}
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}