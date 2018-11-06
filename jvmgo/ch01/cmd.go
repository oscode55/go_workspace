/*
* @Author: myname
* @Date:   2018-08-06 01:27:23
* @Last Modified by:   myname
* @Last Modified time: 2018-08-08 15:34:26
 */
package main

import "flag"
import "fmt"
import "os"

/* os包定义了Args变量,存放传递给命令行的全部参数
   直接处理os.Args变量,需要写很多代码
   所以引入flag包
*/
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage //函数赋值
	//设置需要解析的选项
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse() //解析选项,解析失败则调用 printUsage()函数
	//解析成功调用flag.Args()函数捕获没有被解析的参数
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
