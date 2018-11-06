/*
* @Author: myname
* @Date:   2018-09-12 21:50:51
* @Last Modified by:   myname
* @Last Modified time: 2018-09-12 21:53:37
 */
/*
    //类、超类、接口表中的接口
    CONSTANT_Class_info {
	    u1 tag;
	    u2 name_index;
    }
*/
package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

//根据索引从常量池获取类名
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
