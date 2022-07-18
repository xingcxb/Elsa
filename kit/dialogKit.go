package kit

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// 创建结构体
type aDiaLogKit struct {
}

// DiaLogKit 对外声明一个通用中间方法，用于链式调用
func DiaLogKit() *aDiaLogKit {
	return &aDiaLogKit{}
}

// ErrorDiaLog 错误信息提示
func (d *aDiaLogKit) ErrorDiaLog(content string, window fyne.Window) {
	dialog.ShowError(errors.New(content), window)
}
