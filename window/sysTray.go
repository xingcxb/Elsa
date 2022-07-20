package window

import (
	"Elsa/common"
	"Elsa/kit"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"net/url"
)

// 系统托盘操作

type aTray struct {
}

func SysTray() *aTray {
	return &aTray{}
}

// MakeTray 创建系统托盘
func (t *aTray) MakeTray(a fyne.App) {
	if desk, ok := a.(desktop.App); ok {
		// 软件介绍
		introduceTxt := kit.StrKit().Joint(common.AppName, "  ", common.Version)
		introduceMenu := fyne.NewMenuItem(introduceTxt, nil)
		// 设置托盘按钮列表为禁用状态
		introduceMenu.Disabled = true
		// 托盘列表分割线，此处声明的位置并不重要，重要的是注册位置
		separator := fyne.NewMenuItemSeparator()
		// 划词翻译
		selectMenu := fyne.NewMenuItem("Select Translate", func() {
			fmt.Println("测试...")
			// 测试发送系统通知
			//fyne.CurrentApp().SendNotification(&fyne.Notification{
			//	Title:   "11111",
			//	Content: "22222",
			//})
			// 执行鼠标双击，将数据进行选中
			err := kit.RobotKit().Copy()
			if err != nil {
				fmt.Println(err)
				return
			}
			// 将数据写入到剪切板中
			//_ = kit.RobotKit().WriteSelectionByClipboard("test")
		})
		// 设置划词翻译的快捷键
		selectMenu.Shortcut = &desktop.CustomShortcut{
			KeyName:  fyne.KeyD,
			Modifier: fyne.KeyModifierAlt,
		}

		// 截图翻译
		screenshotsMenu := fyne.NewMenuItem("Screenshots Translate", func() {})
		// 输入翻译
		inputMenu := fyne.NewMenuItem("Input Translate", func() {})
		// 偏好设置
		preferencesMenu := fyne.NewMenuItem("Preferences", func() {})
		// 指向到网站
		goToWebSiteMenu := fyne.NewMenuItem("Author website", func() {
			u, _ := url.Parse(common.WebSite)
			// 通过默认的浏览器打开网址
			_ = a.OpenURL(u)
		})
		menu := fyne.NewMenu("trayMenu", introduceMenu,
			separator, selectMenu, screenshotsMenu, inputMenu,
			separator, preferencesMenu,
			separator, goToWebSiteMenu)
		desk.SetSystemTrayMenu(menu)
	}
}
