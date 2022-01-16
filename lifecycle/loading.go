package lifecycle

import (
	"aliyundriver-webdav/constant"
	"fyne.io/fyne/v2"
	"log"
)

func InitTheme(app fyne.App) {

	resource, err := fyne.LoadResourceFromPath(constant.IconJpgPath())
	if err != nil {
		log.Printf("加载图标失败 %s", err.Error())
	} else {
		app.SetIcon(resource)
	}

}
