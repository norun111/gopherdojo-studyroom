package main

import (
	"flag"
	"fmt"
	"github.com/gopherdojo/gopherdojo-studyroom/kadai1/tomoya-ueno-2907/cvt"
)

func init() {

}

func main() {
	var dir string
	flag.StringVar(&dir ,"dir", "", "dir flag")
	flag.Parse()
	fmt.Println(cvt.DirWalk(dir))
}
