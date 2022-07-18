package kit

import (
	"image"
	"os"
)

type kImageKit struct {
}

func ImageKit() *kImageKit {
	return &kImageKit{}
}

// ReadImg 读取本地图片
func (k *kImageKit) ReadImg(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return img, err
}
