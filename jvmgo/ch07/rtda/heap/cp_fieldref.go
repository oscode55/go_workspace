package heap

/**
  字段引用
*/
import "jvmgo/ch07/classfile"

type FieldRef struct {
	MemberRef
	field *Field //缓存 字段指针
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

// jvms 5.4.3.2
func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	c := self.ResolvedClass() // 先解析类
	//递归 字段 当前类,接口,父类
	field := lookupField(c, self.name, self.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.field = field
}

func lookupField(c *Class, name, descriptor string) *Field {
	//从类解析出来的已经符合的话
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	//从接口表解析
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}
	//从父类解析
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}
