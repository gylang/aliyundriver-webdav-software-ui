package lifecycle

import (
	"aliyundriver-webdav/bussiness"
	"aliyundriver-webdav/m_container"
	"log"
	"time"
)

func MonitorWebdavStatus() {
	timer := time.NewTimer(time.Second * 2)
	SetWebDavStatus()

	go func() {
		for {
			<-timer.C
			SetWebDavStatus()
			get, _ := m_container.MRunningStatus.StatusBinder.Get()
			log.Println("当前webdav状态" + get)
			timer.Reset(time.Second * 2)
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
