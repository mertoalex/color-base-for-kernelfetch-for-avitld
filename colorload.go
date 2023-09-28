package main

import "strings"

func _noerr(unused ...interface{}) {}

var r0="\033[0m"
var c0=strings.Replace(r0,"0m","0;30m",1)
var c1=strings.Replace(r0,"0m","0;31m",1)
var c2=strings.Replace(r0,"0m","0;32m",1)
var c3=strings.Replace(r0,"0m","0;33m",1)
var c4=strings.Replace(r0,"0m","0;34m",1)
var c5=strings.Replace(r0,"0m","0;35m",1)
var c6=strings.Replace(r0,"0m","0;36m",1)
var c7=strings.Replace(r0,"0m","0;37m",1)

func init() {
	_noerr(c0,c1,c2,c3,c4,c5,c6,c7)
}
