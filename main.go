package main

import (
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
	a := app.New()
	window.MainWindows(a)
	//_ = os.Unsetenv("FYNE_FONT")
}
