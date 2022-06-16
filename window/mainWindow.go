package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"tools/common"
)

// MainWindows 创建一个主窗口
func MainWindows(app fyne.App) {
	// 创建一个新窗口
	mainWindow := app.NewWindow(common.AppName)
	// 设置窗口大小
	mainWindow.Resize(fyne.NewSize(500, 500))
	// 设置主体内容
	content(mainWindow)
	//mainWindow.SetContent(widget.NewLabel("test"))
	mainWindow.ShowAndRun()
}

// 设置主体内容
func content(mainWindow fyne.Window) {
	mainWindow.SetContent(widget.NewLabel("测试"))
}
