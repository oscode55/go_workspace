/*
* @Author: myname
* @Date:   2018-09-12 21:10:07
* @Last Modified by:   myname
* @Last Modified time: 2018-09-13 00:40:54
 */
package classfile

// import "fmt"
// import "unicode/utf16"
/*
    CONSTANT_Utf8_info {
	    u1 tag;
	    u2 length;
	    u1 bytes[length];
    }
*/
type ConstantUtf8Info struct {
	str string
}

//读取[]byte 解码成字符串
func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
