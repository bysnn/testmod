package main

import (
	"fmt"

	"github.com/bysnn/testmod/lib"
	"github.com/bysnn/testmod/util"
)

func main() {
	fmt.Println(lib.End("hello, world"))
	fmt.Println(util.Trim(" YING "))

	name := "hello"
	fmt.Println(util.GetSayHi(name))

	maxInt := util.Max(3, 5)
	fmt.Println("maxInt is:", maxInt)
}
