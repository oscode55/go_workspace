/*
* @Author: myname
* @Date:   2018-08-08 15:35:21
* @Last Modified by:   myname
* @Last Modified time: 2018-10-19 17:39:12
*/
package main //特殊包,这个包所在的目录会被编译为可执行文件
import "fmt"
func main(){
	cmd := parseCmd() //解析命令行参数
	if cmd.versionFlag { //用户输入 -version
		printVersion()
	} else if cmd.helpFlag || cmd.class == ""{
		printUsage()
	} else {
		startJVM(cmd)
	}
}
func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n",cmd.cpOption,cmd.class,cmd.args)
}
func printVersion(){
	fmt.Printf("java version \"1.8.0_172\nJava(TM) SE Runtime Environment (build 1.8.0_172-b11)\nJava HotSpot(TM) Client VM (build 25.172-b11, mixed mode, sharing)")
}