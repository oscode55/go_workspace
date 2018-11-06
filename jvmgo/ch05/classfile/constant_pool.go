/*
* @Author: myname
* @Date:   2018-09-12 00:16:41
* @Last Modified by:   myname
* @Last Modified time: 2018-10-17 16:52:56
*/
//常量池结构体
package classfile
type ConstantPool []ConstantInfo

//常量池大小比表头给的数值小1 且从索引1开始 且long double占用两个位置
//读取常量池
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16()) //读取常量池大小
	cp := make([]ConstantInfo,cpCount)
	for i := 1;i<=cpCount-1;i++ { //常量池索引从1开始到,[1,cpCount-1]
		cp[i] = readConstantInfo(reader,cp)
		switch cp[i].(type) {
		case *ConstantLongInfo,*ConstantDoubleInfo:
			i++ //占两个位置
		}
	}
	return cp
}
//按索引查找常量
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index];cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}
//从常量池查找方法的名字和描述符
func (self ConstantPool) getNameAndType(index uint16) (string,string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name,_type

}
//常量池查找类名
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}
//常量池查找UTF-8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}