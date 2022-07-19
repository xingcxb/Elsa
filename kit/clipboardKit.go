package kit

import "github.com/go-vgo/robotgo"

// 剪贴板

type cclipboardKit struct {
}

func ClipboardKit() *cclipboardKit {
	return &cclipboardKit{}
}

// WriteSelectionByClipboard 向剪切板写入数据
func (c *cclipboardKit) WriteSelectionByClipboard(str string) error {
	e := robotgo.WriteAll(str)
	if e != nil {
		return e
	}
	return nil
}
