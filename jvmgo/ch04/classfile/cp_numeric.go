/*
* @Author: myname
* @Date:   2018-09-12 20:21:31
* @Last Modified by:   myname
* @Last Modified time: 2018-09-12 21:13:30
*/
/*
    基本数据类型 数值结构体
    int(byte,short等)、double、float、long
*/
/*
    CONSTANT_Integer_info {
	    u1 tag;
	    u4 bytes;
    }
*/
package classfile
import "math"
//比int小的也放在这个结构体
type ConstantIntegerInfo struct {
	val int32
}
//读取一个uint32数据 转型成int32
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
    bytes := reader.readUint32()
    self.val = int32(bytes)
}
/*
    CONSTANT_Float_info {
	    u1 tag;
	    u4 bytes;
    }
*/
type ConstantFloatInfo struct {
    val float32
}
func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
    bytes := reader.readUint32()
    self.val = math.Float32frombits(bytes)
}
/*
    CONSTANT_Long_info {
	    u1 tag;
	    u4 high_bytes;
	    u4 low_bytes;
    }
*/
type ConstantLongInfo struct {
	val int64
}
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}
/*
    CONSTANT_Double_info {
	    u1 tag;
	    u4 high_bytes;
	    u4 low_bytes;
    }
*/
type ConstantDoubleInfo struct {
	val float64
}
func (self *ConstantDoubleInfo) readInfo (reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
