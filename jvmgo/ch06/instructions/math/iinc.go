/*
* @Author: myname
* @Date:   2018-10-19 14:41:29
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:27:37
 */
package math

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

// Increment local variable by constant
//给局部变量表中的int变量增加常量值
//LB的索引号 常量值 由操作数提供
type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.SetInt(self.Index, val)
}
