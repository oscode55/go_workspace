package heap

import "jvmgo/ch07/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

//类D通过方法符号引用访问类C的某个方法 解析方法
// jvms8 5.4.3.3
func (self *MethodRef) resolveMethodRef() {
	d := self.cp.class
	c := self.ResolvedClass() //先解析类

	if c.IsInterface() { //若C是接口,异常
		panic("java.lang.IncompatibleClassChangeError")
	}

	//根据方法名和描述符解析
	method := lookupMethod(c, self.name, self.descriptor)

	if method == nil { //找不到符合的方法
		panic("java.lang.NoSuchMethodError")
	}
	//类D是否有权限访问该方法
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	//从类以及父类中解析方法
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		//否则从接口 以及父接口解析方法
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
