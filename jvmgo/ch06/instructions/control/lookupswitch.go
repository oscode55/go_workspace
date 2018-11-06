/*
* @Author: myname
* @Date:   2018-10-19 15:36:40
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:23:49
*/
package control

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"

/*
lookupswitch
<0-3 byte pad>
defaultbyte1
defaultbyte2
defaultbyte3
defaultbyte4
npairs1
npairs2
npairs3
npairs4
match-offset pairs...
*/
// Access jump table by key match and jump
type LOOKUP_SWITCH struct {
	defaultOffset int32 //默认偏移
	npairs        int32 //
	//类似字典类型 key是case值 value是偏移指令
	// 索引{0,2,4,6...} 存key
	// 索引{1,3,5,7...} 存value
	matchOffsets  []int32 //容量是 npairs的两倍
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding() //跳出0~3字节
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < self.npairs*2; i += 2 {
		if self.matchOffsets[i] == key {
			offset := self.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(self.defaultOffset))
}
