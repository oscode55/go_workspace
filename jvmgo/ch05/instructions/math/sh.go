/*
* @Author: myname
* @Date:   2018-10-19 14:17:01
* @Last Modified by:   myname
* @Last Modified time: 2018-10-19 14:37:44
*/
package math
import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"
/**
    移位指令 只针对 int long 以及int以下的位数当成int处理
*/
// Shift left int
type ISHL struct {
	base.NoOperandsInstruction
}
//int 只有32位 只取v2前5个比特位就足够表示位移位数了
//go的位移操作符右侧必须是无符号整数
func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt() //先取出的数是位移位数 最大就移动32位
	v1 := stack.PopInt() //后取出的数是变量
	s := uint32(v2) & 0x1f;
	result := v1 << s
	stack.PushInt(result)
}

// Arithmetic shift right int
type ISHR struct{ base.NoOperandsInstruction }

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

// Logical shift right int 逻辑右移
type IUSHR struct{ base.NoOperandsInstruction }
// >>> 运算符 逻辑右移 补0
// 将v1转成无符号位整型 位移后 再转成有符号位
func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32( uint32(v1) >> s )
	stack.PushInt(result)
}

// Shift left long
type LSHL struct{ base.NoOperandsInstruction }
//long有64位 需要取6个比特
func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f // 0011 1111
	result := v1 << s
	stack.PushLong(result)
}

// Arithmetic shift right long
type LSHR struct{ base.NoOperandsInstruction }

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// Logical shift right long
type LUSHR struct{ base.NoOperandsInstruction }

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
