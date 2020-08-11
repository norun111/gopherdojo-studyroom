package cvt

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