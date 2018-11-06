/*
* @Author: myname
* @Date:   2018-08-08 16:16:52
* @Last Modified by:   myname
* @Last Modified time: 2018-08-10 01:53:23
*/
package classpath
import "io/ioutil"
import "path/filepath"
type DirEntry struct {
	absDir string //存放目录的绝对路径
}
func newDirEntry(path string) *DirEntry {
    absDir,err := filepath.Abs(path) //参数转成绝对路径
    if err != nil { //上面的转换有错误
    	panic(err) //终止程序执行
    }
    //创建实例并返回
    return &DirEntry{absDir} 
}
func (self *DirEntry) readClass(className string) ([]byte,Entry,error){
    //把目录和class文件拼成一个完整的路径
    fileName := filepath.Join(self.absDir,className)
    //读取class文件内容
    data,err := ioutil.ReadFile(fileName)
    return data,self,err
}
func (self *DirEntry) String() string {
    return self.absDir //直接返回目录
}