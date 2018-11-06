本章介绍了虚拟机如何加载所需要的类
  1.虚拟机从哪些地方寻找类
    根据类路径class path来搜索类
  2.类路径分为3个部分 
    启动类路径bootstrap classpath
    扩展类路径extendsion classpath
    用户类路径user classpath
  3.启动类默认路径 jre/lib
    拓展类默认路径 jre/lib/ext
    用户类路径 默认为当前路径
  4.启动类或者扩展类 可以使用-Xjre选项作为jre目录
    用户类路径 可以使用-classpath选项
    -classpath/-cp 可覆盖CLASSPATH环境变量
  5.-classpath/-cp 可以指定目录、JAR文件或者ZIP文件
    用法如
    java -cp path\to\classes
    java -cp path\to\lib1.jar
    java -cp path\to\lib2.zip
    java -cp lib\*    //通配符 jdk1.6
    //混合用法
    java -cp path\to\classes;lib\a.jar;lib\c.zip

需要实现的功能
  1.输入的-Xjre 值 主类 [参数]
      解析jre地址[启动类、扩展类路径] 返回 Entry[] //parse
      加载 Entry[] 对主类文件进行解析  //ReadClass
    输入的-classpath/-cp 值 主类 [参数]
      解析cp地址[用户类路径] 返回 Entry[]
      加载 Entry[] 对主类文件进行解析
  2.返回不同的Entry类型 即有不同的readclass函数实现

文件目录结构
  1. -jvmgo
  	-ch01
	-ch02
	    -classpath
	    -cmd.go
	    -main.go
cmd.go文件
  1.添加字段 用来赋值它用户输入命令中jre目录位置
    XjreOption string
  3.parseCmd() 函数
    flag.StringVar(&cmd.XjreOption,"Xjre","","path to jre")

实现类路径
  1.使用组合模式
    类路径由3个小路径组成[启动、扩展、用户类路径]
    三个小路径又由更小的路径组成
  2.定义Entry接口 两个方法
    /*负责寻找和加载class文件
      根据类名称className相对路径 java/lang/Object.class
      返回最终定位到的Entry 读取的字节数 错误信息
    */
    func readClass(className string) ([]byte,Entry,error)

    String() string //返回变量的字符串表示
  3.//根据 路径 的不同创建不同类型的Entry
    //[path有目录、jar、zip、通配符,混合]
    func newEntry(path string) Entry
path为目录DirEntry [entry_dir.go]
  1.DirEntry结构体字段
    absDir //绝对路径
  2.//参数转换成绝对路径
    func newDirEntry(path string) *DirEntry {...}
  3.//目录和class拼完整 读取class文件内容
    func readClass(className string) 
  4.//返回目录
    func String() string
path为ZIP或者JAR文件形式的类路径
  1.ZipEntry 结构体字段
    absPath string //ZIP或JAR文件绝对路径
  2.//参数转成绝对路径
    func newZipEntry(path string) *ZipEntry
  3.//从ZIP文件中提取class文件
    /*
	1.打开ZIP文件
	2.遍历ZIP文件读取class文件
    */
    func (self *ZipEntry) readClass(className string)
path为混合形式的CompositeEntry
  1.//传入的参数其实是路径列表分隔符;隔开
    func newCompositeEntry(pathList string) CompositeEntry
    /*去掉分隔符取一个个path
      生成 newEntry(path) 自动生成所需的Entry类型
      一个个Entry实例存入Entry数组
      返回的类型就是[]Entry{} 数组
    */
  2.func (self CompositeEntry) readClass(className string)
    /*
	依次调用子路径的readClass方法
        for _,entry := range self {
	    data,from,err := entry.readClass(className)
	}
    */
  3.func (self ComposuteEntry) String() string
    //调用每个子路径String()方法 得到的字符串用路径分隔符拼接起来
path为通配符形式
  1./* 
      去*号得到目录路径baseDir 
      遍历目录路径创建ZipEntry 根据后缀名选出JAR文件,跳过子目录
      返回的是compositeEntry 即Entry数组
    */
    func newWildcardEntry(path string) CompositeEntry
使用Entry首先要实现Classpath
  1.ClassPath结构体三个字段 存放三种类路径
    bootClasspath Entry
    extClasspath Entry
    userClasspath Entry
    //解析函数、读取class函数、String字符串表示变量
  2.解析函数
    //使用-Xjre选项解析启动类路径和拓展类路径
    //使用-classpath/-cp解析用户类路径
    //返回Classpath实例
    func Parse(jreOption,cpOption string) *Classpath { ... }
  3.解析启动类路径和拓展类路径
    /*
	1.获取jre路径 如果没有输入-Xjre则在当前目录下找jre目录
  	  若找不到则使用JAVA_HOME环境变量查找
        2.返回混合Entry数组newWildcardEntry赋值给结构体字段
    */
    func (self *Classpath) parseBootAndExtClasspath(jreOption string)
  4.解析用户类路径函数
    /*
	1.用户没有提供-classpath/-cp选项 则使用当前目录作为用户类路径
        2.newEntry(cpOption)根据参数返回实例 赋值给userClassPath
    */
  5.读取class类的函数
    func (self *Classpath) ReadClass(className string)
   //依次从启动类路径、扩展类路径、用户类路径搜索class文件
  6.String()方法返回用户类路径的字符串表示
测试文件main.go
  1.导入classpath目录
  2.重写startJVM函数
    先打印出命令行参数，然后读取主类数据，并打印到控制台