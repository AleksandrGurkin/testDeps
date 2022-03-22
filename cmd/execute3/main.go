package main

import (
	"fmt"

	"github.com/AleksandrGurkin/testDeps/first"
	"github.com/AleksandrGurkin/testDeps/third"
)

var v = 1

func main() {
	fmt.Println("execute1 ", first.First, third.Third)
}
