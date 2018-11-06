/*
* @Author: myname
* @Date:   2018-09-11 00:49:18
* @Last Modified by:   myname
* @Last Modified time: 2018-10-17 16:30:02
 */
package classfile

import "fmt"

type ClassFile struct {
	//magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

//把[]byte解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile, err error) {

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

//依次调用不同的方法解析class文件
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

//检验魔数,是否为一个.class文件
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE { //终止程序进行
		panic("java.lang.ClassFormatError: magic!")
	}
}

//读取和检验版本号jdk1.8
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45: //1.1有次版本号
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 { //1.2开始次版本号都是0
			return
		}
	}
	//版本不是45~52 jdk1.0~1.8
	panic("java.lang.UnsupportedClassVersionError!")
}
func (self *ClassFile) MinorVersion() uint16 { //getter
	return self.minorVersion
}
func (self *ClassFile) MajorVersion() uint16 { //getter
	return self.majorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool { //getter
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 { //getter
	return self.accessFlags
}
func (self *ClassFile) Fields() []*MemberInfo { //getter
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo { //getter
	return self.methods
}

//从常量池查找类名
func (self *ClassFile) ClassName() string {

	return self.constantPool.getClassName(self.thisClass)
}

//从常量池查找超类名
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" //Object没有超类
}
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
