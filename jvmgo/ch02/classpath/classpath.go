/*
* @Author: myname
* @Date:   2018-08-10 02:04:37
* @Last Modified by:   myname
* @Last Modified time: 2018-10-17 15:31:27
*/
package classpath
import "os"
import "path/filepath"
//使用-Xjre选择解析启动类路径和扩展类路径
//使用-classpath/cp 选项解析用户类路径
type Classpath struct {
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}
//解析类 启动类和拓展类 用户类 返回Classpath实例 里面带有3个Entry
func Parse(jreOption,cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption) //-Xjre 解析启动类和扩展类
	cp.parseUserClasspath(cpOption) // -classpath/cp 解析用户类
	return cp
}
//解析启动类和拓展类
func (self *Classpath) parseBootAndExtClasspath(jreOption string){
	jreDir := getJreDir(jreOption) //获取jre目录
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir,"lib","*")
	self.bootClasspath = newWildcardEntry(jreLibPath) //生成启动类的Entry实例
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir,"lib","ext","*")
	self.extClasspath = newWildcardEntry(jreExtPath) //生成拓展类的Entry实例
}
//获取jre目录 没有输入则在当前目录查找
//如果当前目录找不到jre目录 则在系统变量中找
func getJreDir(jreOption string) string {
	//从用户输入的jre路径
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre"){
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh!="" {
		return filepath.Join(jh,"jre")
	}
	panic("Can not find jre folder!")
}
//判断目录是否存在
func exists(path string) bool {
	if _,err := os.Stat(path); err !=nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
//解析用户路径类
func (self *Classpath) parseUserClasspath (cpOption string){
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}
//从启动类、拓展类、用户类中搜索class文件 这个className 不含拓展名
func (self *Classpath) ReadClass(className string)([]byte,Entry,error){
	className = className + ".class"
	if data,entry,err := self.bootClasspath.readClass(className);err == nil{
		return data,entry,err
	}
	if data,entry,err := self.extClasspath.readClass(className);err == nil {
		return data,entry,err
	}
	return self.userClasspath.readClass(className)

}
//返回用户路径类字符串
func (self *Classpath) String() string {
	return self.userClasspath.String()
}