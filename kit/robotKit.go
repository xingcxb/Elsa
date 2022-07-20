package kit

import "github.com/go-vgo/robotgo"

// 系统机器人

type rRobotKit struct {
}

func RobotKit() *rRobotKit {
	return &rRobotKit{}
}

func (r *rRobotKit) Copy() error {
	return robotgo.KeyTap("c", "command")
}

// WriteSelectionByClipboard 向剪切板写入数据
func (c *rRobotKit) WriteSelectionByClipboard(str string) error {
	e := robotgo.WriteAll(str)
	if e != nil {
		return e
	}
	return nil
}
