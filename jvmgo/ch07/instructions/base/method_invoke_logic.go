package base

import "fmt"
import "jvmgo/ch07/rtda"
import "jvmgo/ch07/rtda/heap"

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method) //创建新的帧
	thread.PushFrame(newFrame)          //入jvm stack

	//首先确定方法参数表占用多少位置
	argSlotCount := int(method.ArgSlotCount())

	//n个变量从调用者操作数栈弹出放入
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}

	// hack!
	if method.IsNative() {
		if method.Name() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.Class().Name(), method.Name(), method.Descriptor()))
		}
	}
}
