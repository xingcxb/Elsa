package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	mainWindow.ShowAndRun()
}

// 设置主体内容
func content(mainWindow fyne.Window) {
	text1 := canvas.NewText("1", color.Black)
	text2 := canvas.NewText("2", color.Black)
	text3 := canvas.NewText("3", color.Black)
	grid := container.New(layout.NewGridLayout(2), text1, text2, text3)
	mainWindow.SetContent(grid)
}
