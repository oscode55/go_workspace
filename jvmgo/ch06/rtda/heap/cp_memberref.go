package heap

import "jvmgo/ch06/classfile"

/**
    [基类]成员引用 字段 方法 接口方法
    jvm允许字段[方法]名称相同 但描述符不同
*/
type MemberRef struct {
	SymRef
	name       string //名称
	descriptor string //字段描述符
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

func (self *MemberRef) Name() string {
	return self.name
}
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
