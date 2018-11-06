/*
* @Author: myname
* @Date:   2018-10-19 15:54:34
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:24:34
 */
package extended

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/instructions/loads"
import "jvmgo/ch06/instructions/math"
import "jvmgo/ch06/instructions/stores"
import "jvmgo/ch06/rtda"

/**
 * 注入加载类指令、存储类指令、ret指令和iinc指令
 * 按索引访问局部变量表 索引限制了1字节
 * 如果方法局部变量表范围超过2^8
 * 将无法访问到LB索引
 */
//wide 指令用于改变其他指令的行为
// Extend local variable index by additional bytes
type WIDE struct {
	modifiedInstruction base.Instruction //存放被改变的指令
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8() //读取1字节操作码
	switch opcode {
	//加载指令 拓展成两字节
	case 0x15: //iload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16()) //索引读取了16bit
		self.modifiedInstruction = inst        //指令修改
	case 0x16: //lload
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x17: //fload
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x18: //dload
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x19: //aload
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x36: //istore
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x37: //lstore
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x38: //fstore
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x39: //dstore
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x3a: //astore
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		self.modifiedInstruction = inst
	case 0x84: //iinc 两个操作数 都要扩展成2字节
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		self.modifiedInstruction = inst
	case 0xa9: // ret
		panic("Unsupported opcode: 0xa9!")
	}
}

//只增加索引宽度 而不改变子指令操作
func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}
