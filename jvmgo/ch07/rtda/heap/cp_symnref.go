//[基类]引用
//类 字段 方法 接口方法引用
package heap

// symbolic reference
type SymRef struct {
	cp        *ConstantPool //运行时常量池指针
	className string        //类完全限定名
	class     *Class        //缓存,保证只需解析一次
}

//类符号引用解析
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	//类符号引用以及解析
	return self.class
}

// jvms8 5.4.3.1 类D通过符号引用N引用类C
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) { //检查D是否有权限访问C
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}
