/*
* @Author: myname
* @Date:   2018-10-18 14:34:33
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:22:40
 */
package base

import "jvmgo/ch07/rtda"

//指令的抽象接口
type Instruction interface {
	FetchOperands(reader *BytecodeReader) //字节码中提取操作数
	Execute(frame *rtda.Frame)            //执行指令逻辑
}
type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

//跳转指令
type BranchInstruction struct {
	Offset int //跳转偏移量 两字节
}

//从字节码中提取一个uint16
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

//存储和加载类指令 根据索引存取局部变量表
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint //常量池索引
}

//字节码中读取uint16整数
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}
