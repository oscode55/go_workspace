package references

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

// Create new object
type NEW struct{
	base.Index16Instruction //索引,从常量池找到类符号引用
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()

	//Method Area

	classRef := cp.GetConstant(self.Index).(*heap.ClassRef) //运行时常量池中获取类信息,转为类引用

	class := classRef.ResolvedClass() //解析成运行时类结构体实例
	// todo: init class

	//接口与抽象类不能实例化
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	//Java Heap
	ref := class.NewObject() //创建Object实例 go特性自动赋初值
	//Java Stack
	frame.OperandStack().PushRef(ref) //入栈
}
