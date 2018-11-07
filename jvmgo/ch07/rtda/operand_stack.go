/*
* @Author: myname
* @Date:   2018-10-18 10:42:32
* @Last Modified by:   myname
* @Last Modified time: 2018-10-22 17:32:29
 */
/*
 OperandStack 编译已确定大小
*/
package rtda

import "math"
import "jvmgo/ch07/rtda/heap"

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

//??正确吗
func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}
func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}
func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val >> 32)
	self.size += 2
}
func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)
	return int64(high)<<32 | int64(low)
}
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}
func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}
func (self *OperandStack) PushRef(ref *heap.Object) {
	self.slots[self.size].ref = ref
	self.size++
}
func (self *OperandStack) PopRef() *heap.Object {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil //适应go垃圾回收机制
	return ref
}
func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}

//?? 正确吗
func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}

//返回倒数第i+1个引用
func (self *OperandStack) GetRefFromTop(n uint) *heap.Object {
	return self.slots[self.size-1-n].ref
}

/*
public static float circumference(float r) {
    float pi = 3.14f;
    float area = 2*pi*r
    return area;
}
localvariableTable = 3
stack = 2
局部变量表#0 即r
00 ldc #4    //3.14f入栈
02 fstore_1  //3.14f出栈 放入局部变量表LB编号#1
03 fconst_2  //2.0f入栈
04 fload_1   //LB编号#1 3.14f变量入栈
05 fmul      //栈顶两元素弹出 相乘 2.0f*3.14f 结果入栈
06 fload_0   //LB编号#0入栈
07 fmul      //栈顶两元素弹出 相乘 结果入栈
08 fstore_2  //栈顶元素 存入LB编号#2
09 fload_2   //LB编号#2 入栈
10 return    //栈顶元素出栈 返回给调用者
*/
