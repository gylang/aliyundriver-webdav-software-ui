package m_container

import (
	"aliyundriver-webdav/domain"
	"fyne.io/fyne/v2"
	"sync"
)

var MApp fyne.App
var MWidow fyne.Window

var Config = &domain.AliWebDavConfig{}

var mx = sync.Mutex{}

func DefaultWindow() fyne.Window {
	mx.Lock()
	defer mx.Unlock()
	if MWidow != nil {
		return MWidow
	}
	w := MApp.NewWindow("GiLang-阿里云盘WEBDAV")
	w.SetCloseIntercept(func() {
		w.Close()
	})
	MWidow = w
	return w
}
