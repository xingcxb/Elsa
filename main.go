package main

import (
	"Elsa/common"
	"Elsa/window"
	"fyne.io/fyne/v2/app"
)

func init() {
	// 初始化国际化
	//bundle := i18n.NewBundle(language.English)
	//bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	//bundle.LoadMessageFile("./i18n/en.toml")
	//// 设置中文字体
	//_ = os.Setenv("FYNE_FONT", "./assets/font/HarmonyOS_Sans_SC_Regular.ttf")
}

func main() {
	// 创建一个新的应用，并且设置一个全局唯一的id
	a := app.NewWithID(common.AppId)
	window.MainWindows(a)
	//_ = os.Unsetenv("FYNE_FONT")
}
