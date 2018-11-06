/*
* @Author: myname
* @Date:   2018-09-12 23:20:24
* @Last Modified by:   myname
* @Last Modified time: 2018-09-12 23:23:33
*/
/*
    Deprecated_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
    }
    Synthetic_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
    }
*/
package classfile
type DeprecatedAttribute struct { MarkerAttribute }
type SyntheticAttribute struct { MarkerAttribute }
type MarkerAttribute struct {}
func (self *MarkerAttribute) readInfo(reader *ClassReader){
	//read nothing 这两个属性没有数据
}