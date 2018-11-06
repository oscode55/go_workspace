/*
* @Author: myname
* @Date:   2018-09-12 22:32:10
* @Last Modified by:   myname
* @Last Modified time: 2018-10-17 17:23:09
*/
/*
    属性结构定义
    attribute_info {
	    u2 attribute_name_index;
	    u4 attribute_length;
	    u1 info[attribute_length];
    }
*/
package classfile
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}
//读取属性表
func readAttributes(reader *ClassReader,cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make( []AttributeInfo,attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader,cp)
	}
	return attributes
}
//读取单个属性
func readAttribute(reader *ClassReader,cp ConstantPool) AttributeInfo {
    attrNameIndex := reader.readUint16() //属性名索引
    attrName := cp.getUtf8(attrNameIndex)
    attrLen := reader.readUint32() //属性长度
    attrInfo := newAttributeInfo(attrName,attrLen,cp) //根据名称,创建属性实例
    attrInfo.readInfo(reader) //不同属性 具有不同的读法
    return attrInfo
}
func newAttributeInfo(attrName string,attrLen uint32,cp ConstantPool) AttributeInfo {
    switch attrName {
    	case "Code": return &CodeAttribute{cp :cp}
    	case "ConstantValue": return &ConstantValueAttribute{}
    	case "Deprecated": return &DeprecatedAttribute{}
    	case "Exceptions": return & ExceptionsAttribute{}
    	case "LineNumberTable": return &LineNumberTableAttribute{}
    	case "LocalVariableTable": return &LocalVariableTableAttribute{}
    	case "SourceFile": return &SourceFileAttribute{cp:cp}
    	case "Synthetic": return &SyntheticAttribute{}
    	default: return &UnparsedAttribute{attrName,attrLen,nil}
    }
}