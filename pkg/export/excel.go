package export

import "gin_log/pkg/setting"

func GetExcelFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
}
func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}
func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + setting.AppSetting.ExportSavePath
}
