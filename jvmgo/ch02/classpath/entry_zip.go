/*
* @Author: myname
* @Date:   2018-08-08 16:24:37
* @Last Modified by:   myname
* @Last Modified time: 2018-08-08 16:36:24
*/
package classpath
import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"
type ZipEntry struct {
	absPath string //存放zip或jar文件的绝对路径
}
func newZipEntry(path string) *ZipEntry{
	absPath,err :=filepath.Abs(path)
	if err !=nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}
//从zip文件中提取class文件
func (self *ZipEntry) readClass(className string)([]byte,Entry,error){
    //打开zip文件 
    r,err := zip.OpenReader(self.absPath)
    //出错直接返回
    if err != nil{
    	return nil,nil,err
    }
    defer r.Close()
    //遍历包里的文件
    for _,f:=range r.File {
    	//看能否找到class文件
    	if f.Name == className {
    		//文件打开
    		rc,err := f.Open()
    		if err != nil {
    			return nil,nil,err
    		}
    		defer rc.Close()
    		//读取文件数据
    		data,err := ioutil.ReadAll(rc)
    		if err!=nil {
    			return nil,nil,err
    		}
    		return data,self,nil
    	}
    }
    return nil,nil,errors.New("class not found: "+className)
}
func (self *ZipEntry) String() string {
	return self.absPath
}