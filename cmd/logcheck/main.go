package main

import (
	"fmt"
	"github.com/LonelyPale/logcheck"
)

func main() {
	fmt.Println("hello world")
	n, _ := logcheck.CountLine(nil)
	fmt.Println(n)
}
