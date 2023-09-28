package main

import (
	"fmt"
	"strings"
)
//this file for functions to get system info

func getBanner() string {
	return fmt.Sprintf(strings.Replace(`SOF
%s   ___  ___   _  ___  _________	%s
%s  / _ )/ _ | / |/ / |/ / __/ _ \	%s
%s / _  / __ |/    /    / _// , _/	%s
%s/____/_/ |_/_/|_/_/|_/___/_/|_|	%s
%s--------------------------------	%s`,"SOF\n","",1),
	c1,r0,c3,r0,c2,r0,c4,r0,c5,r0)
}

func getKernel() string {
	return "Linux"
}
