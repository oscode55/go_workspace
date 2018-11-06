/*
* @Author: myname
* @Date:   2018-09-13 00:19:43
* @Last Modified by:   myname
* @Last Modified time: 2018-09-13 00:36:05
*/
/*
    LineNumberTable_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	    u2 line_number_table_length;
	    {
		    u2 start_pc;
		    u2 line_number;
	    }  line_number_table[line_number_table_length];
    }
*/
package classfile
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}
type LineNumberTableEntry struct {
	startPc uint16
	lineNumber uint16
}
//读取属性表数据
func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
    lineNumberTableLength := reader.readUint16()
    self.lineNumberTable = make([]*LineNumberTableEntry,lineNumberTableLength)
    for i := range self.lineNumberTable {
    	self.lineNumberTable[i] = &LineNumberTableEntry {
    		startPc: reader.readUint16(),
    		lineNumber: reader.readUint16(),
    	}
    }
}
