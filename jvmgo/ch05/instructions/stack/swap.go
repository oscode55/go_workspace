/*
* @Author: myname
* @Date:   2018-10-19 13:46:30
* @Last Modified by:   myname
* @Last Modified time: 2018-10-19 16:54:32
 */
package stack

import "jvmgo/ch05/instructions/base"
import "jvmgo/ch05/rtda"

// Swap the top two operand stack values
type SWAP struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
          \/
          /\
         V  V
[...][c][a][b]
*/
func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
