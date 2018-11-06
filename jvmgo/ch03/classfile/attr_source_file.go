/*
* @Author: myname
* @Date:   2018-09-12 23:24:47
* @Last Modified by:   myname
* @Last Modified time: 2018-09-12 23:27:09
*/
/*
    SourceFile_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	    u2 sourcefile_index;
    }
*/
package classfile
type SourceFileAttribute struct {
	cp ConstantPool
	sourceFileIndex uint16
}
func (self *SourceFileAttribute) readInfo(reader *ClassReader){
	self.sourceFileIndex = reader.readUint16()
}
func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}