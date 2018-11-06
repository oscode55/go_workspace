/*
* @Author: myname
* @Date:   2018-08-08 15:35:21
* @Last Modified by:   myname
* @Last Modified time: 2018-10-19 16:46:30
*/
package main //特殊包,这个包所在的目录会被编译为可执行文件
import "fmt"
import "strings"
import "jvmgo/ch05/classfile"
import "jvmgo/ch05/classpath"
func main(){
	cmd := parseCmd() //解析命令行参数
	if cmd.versionFlag { //用户输入 -version
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == ""{
		printUsage()
	} else {
		startJVM(cmd)
	}
}
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption,cmd.cpOption)
	className := strings.Replace(cmd.class,".","/",-1)
    cf := loadClass(className,cp)
    mainMethod := getMainMethod(cf)
    if mainMethod != nil {
    	interpreter(mainMethod)
    }else {
    	fmt.Printf("Main method not found in class %s\n", cmd.class)
    }
}
func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _,m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
//读取并解析class文件
func loadClass(className string,cp *classpath.Classpath) *classfile.ClassFile {
	classData,_,err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf,err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}
func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n",cf.MajorVersion(),cf.MinorVersion())
	fmt.Printf("constants count: %v\n",len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n",cf.AccessFlags())
	fmt.Printf("this class: %v\n",cf.ClassName())
	fmt.Printf("super class: %v\n",cf.SuperClassName())
	fmt.Printf("interfaces: %v\n",cf.InterfaceNames())
	fmt.Printf("fields count: %v\n",len(cf.Fields()))
	for _,f := range cf.Fields() {
		fmt.Printf(" %s\n",f.Name())
	}
	fmt.Printf("methods count: %v\n",len(cf.Methods()))
	for _,m := range cf.Methods() {
		fmt.Printf(" %s\n",m.Name())
	}
}