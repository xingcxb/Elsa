package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"tools/common"
)

// MainWindows 创建一个主窗口
func MainWindows(app fyne.App) {
	// 创建一个新窗口
	mainWindow := app.NewWindow(common.AppName)
	// 设置窗口大小
	mainWindow.Resize(fyne.NewSize(1000, 700))
	// 设置主体内容
	content(mainWindow)
	// 窗口的显示和运行
	mainWindow.ShowAndRun()
}

// 设置主体内容
func content(mainWindow fyne.Window) {
	text1 := canvas.NewText("111111111111", color.Black)
	text2 := canvas.NewText("222222222222", color.Black)
	text3 := canvas.NewText("333333333333", color.Black)
	grid := container.NewHBox(text1, text2, text3)
	mainWindow.SetContent(grid)
}
