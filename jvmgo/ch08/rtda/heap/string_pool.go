package heap

import "unicode/utf16"

var internedStrings = map[string]*Object{}

// todo
// go string -> java.lang.String
//go字符串 返回 java类型的字符串对象
func JString(loader *ClassLoader, goStr string) *Object {
	//如果已经在池中
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	chars := stringToUtf16(goStr) //将go的UTF8转成java的UTF16字符数组
	//根据UTF16格式创建的字符数组 转成java的字符数组对象
	jChars := &Object{loader.LoadClass("[C"), chars}
	//根据java字符数组对象 生成String对象
	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)

	internedStrings[goStr] = jStr
	return jStr
}

// java.lang.String -> go string
func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

// utf8 -> utf16
func stringToUtf16(s string) []uint16 {
	runes := []rune(s)         // utf32
	return utf16.Encode(runes) // func Encode(s []rune) []uint16
}

// utf16 -> utf8
func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	return string(runes)
}
