package kit

import "strings"

type kStrKit struct {
}

func StrKit() *kStrKit {
	return &kStrKit{}
}

// Joint 拼接字符串
func (k *kStrKit) Joint(strs ...string) string {
	var result strings.Builder
	for _, str := range strs {
		result.WriteString(str)
	}
	return result.String()
}
