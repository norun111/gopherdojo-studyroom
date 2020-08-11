package cvt

import (
	"flag"
	"fmt"
	"os"
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

	arg1 := os.Args[1]
	fs.Parse(os.Args[2:])

	fmt.Println("dir:", arg1)
	fmt.Println("arg1:", *bx)
	fmt.Println("arg2:", *ax)

	return nil
}

