/*
* @Author: myname
* @Date:   2018-09-13 00:13:54
* @Last Modified by:   myname
* @Last Modified time: 2018-09-13 00:19:29
 */
/*
    //异常属性
    Exceptions_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	    u2 number_of_exceptions;
	    u2 exception_index_table[number_of_exceptions];
    }
*/
package classfile

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}
func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
