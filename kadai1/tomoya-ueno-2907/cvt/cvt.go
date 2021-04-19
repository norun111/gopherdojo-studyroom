package cvt

import (
	"io/ioutil"
	"path/filepath"
)

func DirWalk(dir string) []string {
	files, _ := ioutil.ReadDir(dir)

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, DirWalk(filepath.Join(dir, file.Name()))...)
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths
}
