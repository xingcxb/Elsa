package main

import (
	"fyne.io/fyne/v2/app"
	"os"
	"tools/window"
)

func init() {
	// 设置中文字体
	_ = os.Setenv("FYNE_FONT", "./assets/font/HarmonyOS_Sans_SC_Regular.ttf")

}

func main() {
	a := app.New()
	window.MainWindows(a)
	_ = os.Unsetenv("FYNE_FONT")
}
