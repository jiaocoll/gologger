package main

import (
	"gologger/pkg/format"
	"gologger/pkg/level"
)

/*
创建人员：云深不知处
创建时间：2022/2/3
程序功能：主程序
*/

func Gologinit(level level.Thelevel) *format.Logger {
	L := &format.Logger{
		Level: level,
	}
	return L
}

func main() {
	l := Gologinit(1)
	l.VEBO("[ICMPCheck]:", "OK!")
}
