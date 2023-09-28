package main

import "fmt"

func assert(err bool, msg ...interface{}) {
        if (!err) {
                error(msg...)
                panic(err)
        }
}

func error(msg ...interface{}) {
        fmt.Printf("%s[E]: %s", string(c1), string(r0))
        fmt.Println(msg...)
}
