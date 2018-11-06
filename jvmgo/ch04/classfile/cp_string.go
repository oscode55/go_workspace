/*
* @Author: myname
* @Date:   2018-09-12 21:43:52
* @Last Modified by:   myname
* @Last Modified time: 2018-09-12 21:48:07
*/
/*
//java.lang.String 字面量
    CONSTANT_String_info {
	    u1 tag;
	    u2 string_index; //常量池索引
    }
*/
package classfile
type ConstantStringInfo struct {
	cp ConstantPool
	stringIndex uint16
}
//读取常量池索引
func (self *ConstantStringInfo) readInfo(reader *ClassReader){
	self.stringIndex = reader.readUint16()
}
//按索引从常量池中查找字符串
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}