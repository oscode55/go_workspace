/*
* @Author: myname
* @Date:   2018-08-08 15:35:21
* @Last Modified by:   myname
* @Last Modified time: 2018-08-10 02:30:54
 */
package main //特殊包,这个包所在的目录会被编译为可执行文件
import "fmt"
import "strings"
import "jvmgo/ch02/classpath"

func main() {
	cmd := parseCmd()    //解析命令行参数
	if cmd.versionFlag { //用户输入 -version
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	//打印出命令行参数
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	//读取主类信息
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
