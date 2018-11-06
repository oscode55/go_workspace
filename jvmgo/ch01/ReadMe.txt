本章介绍了go语言的基本目录结构
并实现简单的命令行输入执行工具
  1.src go语言源代码 (必须要有)
  2.pkg 编译好的包对象文件 (可自动创建)
  3.bin 链接好的可执行文件 (可自动创建)
我的工作目录结构
  1. D:\go_workspace
  2.系统变量
    GOPATH=D:\go_workspace
    PATH+=D:\go\bin //go.exe添加到全局
  3.第一章文件
    D:\go_workspace\src
        -jvmgo
 	    -ch01
  4.编译项目下的所有源代码
    go install jvmgo\ch01 //结果是生成 bin\ch01.exe
本章主要搭建一个简单的命令行输入程序
  1.必须实现至少下列三种 使用GO语言内置的很多方便函数
  java -version
  java -help
  java [-options] class [args]
  //-help 或者未输入class名称 则输出提示
  //-version 打印版本号
  //否则 开启jvm函数 即打印classpath和args信息