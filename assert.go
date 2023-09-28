package main

import "fmt"

func assert(err bool, msg ...interface{}) {
        if (!err) {
                error(msg...)
                panic(err)
        }
}

func error(msg ...interface{}) {
        colorReset := "\033[0m"
        colorRed := "\033[31m"
        fmt.Printf("%s[E]: %s", string(colorRed), string(colorReset))
        fmt.Println(msg...)
}
