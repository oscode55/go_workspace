/*
* @Author: myname
* @Date:   2018-09-12 23:16:13
* @Last Modified by:   myname
* @Last Modified time: 2018-09-12 23:17:33
*/
package classfile
type UnparsedAttribute struct {
	name string
	length uint32
	info []byte
}
func (self *UnparsedAttribute) readInfo(reader *ClassReader){
	self.info = reader.readBytes(self.length)
}