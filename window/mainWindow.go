package window

import (
	"Elsa/common"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
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
	//mainWindow.ShowAndRun()
	app.Lifecycle()
	// 运行不展示窗口
	app.Run()
}

// 设置主体内容
func content(mainWindow fyne.Window) {
	leftTabMenu(mainWindow)
	//mainWindow.SetContent(grid)
}

// 左侧目录
func leftTabMenu(mainWindow fyne.Window) {
	// 翻译
	translateImg := canvas.NewImageFromFile(common.TranslateImage)
	translateImg.FillMode = canvas.ImageFillOriginal
	translateImg.Resize(fyne.NewSize(40, 40))
	translateTxt := canvas.NewText("translate", color.Black)

	// ORC
	orcImg := canvas.NewImageFromFile(common.OrcImage)
	orcImg.FillMode = canvas.ImageFillOriginal
	orcImg.Resize(fyne.NewSize(40, 40))
	orcTxt := canvas.NewText("orc", color.Black)

	// 通用
	generalImg := canvas.NewImageFromFile(common.GeneralImage)
	generalImg.FillMode = canvas.ImageFillOriginal
	generalImg.Resize(fyne.NewSize(40, 40))
	generalTxt := canvas.NewText("general", color.Black)

	coreMenu := container.NewVBox(translateImg, translateTxt, orcImg, orcTxt, generalImg, generalTxt)
	coreMenu.Resize(fyne.NewSize(80, 700))

	mainWindow.SetContent(coreMenu)
}
