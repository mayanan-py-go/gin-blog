package file

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}
func GetExt(fileName string) string {
	return path.Ext(fileName)
}
func CheckExist(src string) bool {
	_, err := os.Stat(src)
	return errors.Is(err, fs.ErrExist)
}
func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return errors.Is(err, fs.ErrPermission)
}
func IsNotExistMkDir(src string) error {
	exist := CheckExist(src)
	if !exist {
		return Mkdir(src)
	}
	return nil
}
func Mkdir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	return err
}
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}
