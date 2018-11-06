/*
* @Author: myname
* @Date:   2018-09-12 23:29:32
* @Last Modified by:   myname
* @Last Modified time: 2018-10-17 17:30:05
 */
/*
    表示常量表达式的值
    ConstantValue_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;//必须是2
	    u2 constantvalue_index;
    }
*/
package classfile

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}
func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
