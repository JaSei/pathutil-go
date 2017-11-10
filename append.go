package pathutil

import (
	"github.com/pkg/errors"
)

func (path pathImpl) Append(data string) error {
	return path.AppendBytes([]byte(data))
}

func (path pathImpl) AppendBytes(data []byte) (err error) {
	file, err := path.OpenWriterAppend()
	if err != nil {
		return errors.Wrap(err, "Append open failed")
	}
	defer func() {
		if e := file.Close(); e != nil {
			err = errors.Wrap(e, "Append close failed")
		}
	}()

	_, err = file.Write(data)

	return errors.Wrap(err, "Append write failed")
}
