package qrcode

import (
	"gin_log/pkg/file"
	"gin_log/pkg/setting"
	"gin_log/pkg/util"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/jpeg"
	"os"
)

type QrCode struct {
	URL string
	Width int
	Height int
	Ext string
	Level qr.ErrorCorrectionLevel
	Mode qr.Encoding
}
const (
	EXT_JPG = ".jpg"
)
func NewQrCode(url string, width, height int, ext string, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL: url,
		Width: width,
		Height: height,
		Ext: ext,
		Level: level,
		Mode: mode,
	}
}

func GetQrCodePath() string {
	return setting.AppSetting.QrCodeSavePath
}
func GetQrCodeFullPath() string {
	return setting.AppSetting.RuntimeRootPath + setting.AppSetting.QrCodeSavePath
}
func GetQrCodeFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetQrCodePath() + name
}
func GetQrCodeFileName(value string) string {
	return util.EncodeMD5(value)
}
func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}
func (q *QrCode) CheckEncode(path string) bool {
	src := path + GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	return !file.CheckExist(src)
}
func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name

	err := file.IsNotExistMkDir(path)
	if err != nil {
		return "", "", err
	}

	if !file.CheckExist(src) {
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}
		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}
		f, err := file.Open(src, os.O_CREATE|os.O_RDWR, 0755)
		if err != nil {
			return "", "", err
		}
		defer func() {
			_ = f.Close()
		}()
		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}
	}
	return name, path, nil
}
