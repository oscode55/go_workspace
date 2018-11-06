/*
* @Author: myname
* @Date:   2018-10-19 16:15:34
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:24:23
 */
package extended

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

// Branch always (wide index)
//在 goto 指令的基础上 索引从2字节 变成4字节
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}
func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
