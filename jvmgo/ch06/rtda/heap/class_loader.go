package heap

import "fmt"
import "jvmgo/ch06/classfile"
import "jvmgo/ch06/classpath"
/**
    类加载器
 */
/*
class names:
    - primitive types: boolean, byte, int ...
    - primitive arrays: [Z, [B, [I ...
    - non-array classes: java/lang/Object ...
    - array classes: [Ljava/lang/Object; ...
*/
type ClassLoader struct {
	cp       *classpath.Classpath //类路径指针
	classMap map[string]*Class // loaded classes 已经加载的类
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}
//类数据加载到方法区
func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		// already loaded
		return class
	}

	return self.loadNonArrayClass(name)
}
//非数组加载 数组由JVM运行期生成 而非来自class
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name) //数据读入内存
	class := self.defineClass(data) //解析class 生成可用运行类数据[方法区]
	link(class) //链接
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}
//为了打印类加载信息,类路径也返回给调用者
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// jvms 5.3.5
func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data) //静态class数据转换成运行时类数据
	class.loader = self

	resolveSuperClass(class) //加载父类
	resolveInterfaces(class) //加载接口表

	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		//panic("java.lang.ClassFormatError")
		panic(err)
	}
	return newClass(cf)
}

// jvms 5.4.3.1 非Object递归加载父类
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	// todo
}

// jvms 5.4.2 准备阶段
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)//实例字段个数 编号
	calcStaticFieldSlotIds(class) //静态字段个数 编号
	allocAndInitStaticVars(class) //为静态成员 静态常量赋零值
}
//计算实例字段个数 同时编号
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	//实例字段个数 要递归计算父类的
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}
//计算静态字段个数 同时编号
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}
//给类变量分配空间 赋初始值
func allocAndInitStaticVars(class *Class) {
	//静态成员 go的特性保证其符合零值
	class.staticVars = newSlots(class.staticSlotCount)
	//静态常量成员
	for _, field := range class.fields {
		//静态 常量 编译器可知 存储在常量池
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;": //字符串常量
			panic("todo")
		}
	}
}
