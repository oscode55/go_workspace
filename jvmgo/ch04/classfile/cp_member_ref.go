/*
* @Author: myname
* @Date:   2018-09-12 22:04:54
* @Last Modified by:   myname
* @Last Modified time: 2018-09-13 12:04:10
 */
/*
    字段符号引用/普通方法符号引用/接口方法符号引用
    CONSTANT_Fieldref_info {
	    u1 tag;
	    u2 class_index;
	    u2 name_and_type_index;
    }
*/
package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}
func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

//字段符号引用 方法符号引用 接口方法符号引用
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }
