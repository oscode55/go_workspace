package references

/**
  putstatic x1 x2
  操作数1: uint16索引 从常量池找到字段符号引用 并解析
  操作数2: 要给静态变量赋的值 从操作数栈中弹出
*/
import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtda"
import "jvmgo/ch06/rtda/heap"

// Set static field in class
type PUT_STATIC struct{ base.Index16Instruction }

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()       //当前方法
	currentClass := currentMethod.Class() //当前类
	cp := currentClass.ConstantPool()     //常量池指针
	//获取字段符号引用
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	//解析字段,指向真实运行时数据结构
	field := fieldRef.ResolvedField()
	//获取类
	class := field.Class()
	// todo: init class

	//不能是非静态 或者是常量
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	//为字段赋值 值从操作数栈中弹出
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	default:
		// todo
	}
}
