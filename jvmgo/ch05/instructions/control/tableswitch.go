/*
* @Author: myname
* @Date:   2018-10-19 15:16:26
* @Last Modified by:   myname
* @Last Modified time: 2018-10-19 15:32:01
 */
/**
 * 如果case值可以编码成一个索引表 实现成tableswitch
 * 否则实现成 lookupswitch指令
 */
package control

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

/*
tableswitch
<0-3 byte pad> //0~3填充 以保证defaultOffset地址是4的倍数
defaultbyte1   //对应执行跳转所需的字节码偏移量
defaultbyte2
defaultbyte3
defaultbyte4
lowbyte1    //记录case的取值范围
lowbyte2
lowbyte3
lowbyte4
highbyte1
highbyte2
highbyte3
highbyte4   //记录case的取值范围
jump offsets... //索引表 high-low+1 个int值 对应各种情况下 跳转所需的偏移量
*/
// Access jump table by index and jump
type TABLE_SWITCH struct {
	defaultOffset int32   //u4 默认偏移量
	low           int32   //u4 偏移量起始
	high          int32   //u4 偏移量结束
	jumpOffsets   []int32 //索引表 对应每种情况对应的偏移量 长度[high-low+1]
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()                    //0~3填充 以保证defaultOffset地址是4的倍数
	self.defaultOffset = reader.ReadInt32() //对应执行跳转所需的字节码偏移量
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {

	index := frame.OperandStack().PopInt() //操作数栈中弹出一个int变量

	//查看是否在[low,high]之间 如果在根据索引表 否则按defaultOffset
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}

	base.Branch(frame, offset)
}
