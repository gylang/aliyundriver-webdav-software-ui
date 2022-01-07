package lifecycle

import (
	"aliyundriver-webdav/bussiness"
	"aliyundriver-webdav/m_container"
	"fyne.io/fyne/v2/widget"
	"log"
	"time"
)

func MonitorWebdavStatus(w *widget.Label) {
	timer := time.NewTimer(time.Second * 2)
	SetWebDavStatus()

	go func() {
		for {
			<-timer.C
			SetWebDavStatus()
		}
	}()
}

func SetWebDavStatus() {
	process := bussiness.ListProcess("IMAGENAME eq aliyundrive-webdav.exe")
	if len(process) > 0 {
		m_container.MRunningStatus.Running = true
		err := m_container.MRunningStatus.StatusBinder.Set("运行中")
		if err != nil {
			log.Fatal()
		}

	} else {
		m_container.MRunningStatus.Running = false
		err := m_container.MRunningStatus.StatusBinder.Set("未启动")
		if err != nil {
			log.Fatal()
		}
	}
}
