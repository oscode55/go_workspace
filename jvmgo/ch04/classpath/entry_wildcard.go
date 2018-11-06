/*
* @Author: myname
* @Date:   2018-08-08 16:50:04
* @Last Modified by:   myname
* @Last Modified time: 2018-08-10 02:39:05
 */
package classpath

import "os"
import "path/filepath"
import "strings"

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] //remove * 得到baseDir
	compositeEntry := []Entry{}
	//根据后缀名选出JAR文件
	//返回SkipDir跳过子目录,通配符类路径不能递归匹配子目录下的JAR文件
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	//遍历baseDir创建ZipEntry
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
