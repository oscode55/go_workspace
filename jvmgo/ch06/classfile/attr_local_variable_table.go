/*
* @Author: myname
* @Date:   2018-09-13 00:52:39
* @Last Modified by:   myname
* @Last Modified time: 2018-09-13 11:20:07
 */
/*
    LocalVariableTable_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	    u2 local_variable_table_length;
	    {
	 	    u2 start_pc;
	 	    u2 length;
	 	    u2 name_index;
	 	    u2 descriptor_index;
	 	    u2 index;
	    } local_varibale_table[ local_varibale_table_length ]
    }

*/
package classfile

type LocalVariableTableAttribute struct {
	cp                 ConstantPool
	localVariableTable []*LocalVariableTableEntry
}
type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

//读取属性表数据
func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVaribaleTableLength := reader.readUint16()
	self.localVariableTable = make([]*LocalVariableTableEntry, localVaribaleTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}
