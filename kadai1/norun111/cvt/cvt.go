package cvt

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)



type Cvt struct {
	RootDir string
	BefExt string
	AftExt string
}

func New() *Cvt {
	return &Cvt{
		RootDir: ".",
		BefExt: "jpg",
		AftExt: "png",
	}
}

func (cvt *Cvt) Run() error {
	cvt.SetFlag()

	return nil
}

func (cvt *Cvt) SetFlag() error {

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	var (
		bx = fs.String("b", "jpeg", "file extension before executing")
		ax = fs.String("a", "png", "file extension after executing")
	)

	arg1 := os.Args[1] // .
	fs.Parse(os.Args[2:])

	fmt.Println("dir:", arg1) //これをEncodeFileに渡す
	fmt.Println("arg1:", *bx)
	fmt.Println("arg2:", *ax)

	//filepath := dirWalk(arg1)
	//cvt.EncodeFile(filepath)

	return nil
}

func dirWalk(dir string) []string {
	images, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string

	for _, image := range images {
		if image.IsDir() {
			paths = append(paths, dirWalk(filepath.Join(dir, image.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, image.Name()))
	}
	return paths
}

func (cvt *Cvt) EncodeFile(filepath []string) error {
	fmt.Println(filepath)

	return nil
}

