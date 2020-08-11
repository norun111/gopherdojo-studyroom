package main

import (
	"./cvt"
	"fmt"
)

func main() {
	cli := cvt.New()
	fmt.Println(cli.Run())
}

