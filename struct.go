package utilpath

import (
	"os"
)

type Path struct {
	Path string
	file *os.File
}

type TempFileOpt struct {
	Dir    string
	Prefix string
}
