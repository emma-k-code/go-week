package phone

import (
	"fmt"
)

type Phone interface {
	GetName() string
	Camera() string
	Microphone() string
}

func ClickShutter(p Phone) {
	fmt.Printf("歡迎使用 %s \n", p.GetName())
	fmt.Printf("使用 %s 拍照 \n", p.Camera())
}

func Recording(p Phone) {
	fmt.Printf("使用 %s 錄音 \n", p.Microphone())
}
