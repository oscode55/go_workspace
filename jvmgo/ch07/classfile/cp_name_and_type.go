/*
* @Author: myname
* @Date:   2018-09-12 21:56:12
* @Last Modified by:   myname
* @Last Modified time: 2018-09-13 12:12:53
 */
/*
    字段或方法 的名称和描述符
    CONSTANT_NameAndType_info {
	    u1 tag;
	    u2 name_index;
	    u2 descriptor_index;
    }
*/
package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
