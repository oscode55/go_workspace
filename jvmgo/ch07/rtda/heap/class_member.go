/**
  类成员变量 字段或者方法
*/
package heap

import "jvmgo/ch07/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class //Class结构体指针
}

//从class文件中复制数据
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

// getters
func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}

// jvms 5.4.4
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	//public方法
	if self.IsPublic() {
		return true
	}
	c := self.class
	//protected 需要同包 或子类
	if self.IsProtected() {
		return d == c || d.isSubClassOf(c) ||
			c.GetPackageName() == d.GetPackageName()
	}
	//default 缺省必须同包
	if !self.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}
	//private类型 只有声明这个字段的类才能访问
	return d == c
}
