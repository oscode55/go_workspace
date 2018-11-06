/*
* @Author: myname
* @Date:   2018-08-08 16:36:50
* @Last Modified by:   myname
* @Last Modified time: 2018-08-10 02:37:30
 */
package classpath

import "errors"
import "strings"

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	//数组属于低层数据结构,一般用slice类型
	//按分隔符分成小路径 把小路径转成Entry实例
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	//依次调用子路径的readClass方法
	for _, entry := range self {
		//成功读取返回数据
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	//遍历完还没有找到 则返回错误信息
	return nil, nil, errors.New("class not found: " + className)
}
func (self CompositeEntry) String() string {
	//调用每个子路径的String()方法
	//把得到的字符串用路径分隔符拼接
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
