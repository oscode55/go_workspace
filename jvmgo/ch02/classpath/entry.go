/*
* @Author: myname
* @Date:   2018-08-08 16:06:31
* @Last Modified by:   myname
* @Last Modified time: 2018-08-10 02:33:17
*/
package classpath
import "os"
import "strings"
const pathListSeparator = string(os.PathListSeparator)//存放路径分割符
type Entry interface {
	//寻找和加载class路径文件,参数是class的相对路径 如java/lang/Object.class
	//返回值是读取到的字节数、最终定位到class文件的Entry和错误信息error
	readClass(className string)([]byte,Entry,error)
	//返回变量的字符串表示
	String() string
}
//根据参数创建不同类型的Entry实例
func newEntry(path string) Entry {
	if strings.Contains(path,pathListSeparator){
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path,"*"){
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR") ||
	    strings.HasSuffix(path,".zip") || strings.HasSuffix(path,".ZIP"){
    	return newZipEntry(path);
	}
	return newDirEntry(path);//目录形式类路径
}