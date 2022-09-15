package cache_service

import (
	"errors"
	"gin_log/models"
	"gin_log/pkg/e"
	"gin_log/pkg/export"
	"gin_log/pkg/file"
	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize/v2"
	"io"
	"strconv"
	"strings"
	time2 "time"
)

type Tag struct {
	ID int
	Name string
	State int

	PageNum int
	PageSize int
}
func (t *Tag) GetTagsKey() string {
	keys := []string{e.CACHE_TAG, "LIST"}
	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	if t.State >= 0 {
		keys = append(keys, strconv.Itoa(t.State))
	}
	if t.PageNum > 0 {
		keys = append(keys, strconv.Itoa(t.PageNum))
	}
	if t.PageSize > 0 {
		keys = append(keys, strconv.Itoa(t.PageSize))
	}
	return strings.Join(keys, "_")
}
func (t *Tag) Export() (string, error) {
	tags := models.GetTags(t.PageNum, t.PageSize, map[string]any{"name": t.Name, "state": t.State})
	if len(tags) == 0 {
		return "", errors.New("暂无导出数据")
	}

	f := xlsx.NewFile()
	sheet, err := f.AddSheet("标签信息")
	if err != nil {
		return "", err
	}

	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	for _, tag := range tags {
		values := []string{
			strconv.Itoa(tag.ID),
			tag.Name,
			tag.CreatedBy,
			strconv.Itoa(int(tag.CreatedOn)),
			tag.ModifiedBy,
			strconv.Itoa(int(tag.ModifiedOn)),
		}
		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	time := strconv.Itoa(int(time2.Now().Unix()))
	filename := "tag-" + time + ".xlsx"

	fullPath := export.GetExcelFullPath() + filename
	err = file.IsNotExistMkDir(export.GetExcelFullPath())
	if err != nil {
		return "", err
	}
	err = f.Save(fullPath)
	if err != nil {
		return "", err
	}
	return filename, nil
}
func (t *Tag) Import(r io.Reader) error {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}
	rows, err := f.GetRows("标签信息")
	if err != nil {
		return err
	}
	for iRow, row := range rows {
		if iRow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}
			models.AddTag(data[1], 1, data[2])
		}
	}
	return nil
}
