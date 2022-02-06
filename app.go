package main

import (
	"aliyundriver-webdav/bussiness"
	"aliyundriver-webdav/lifecycle"
	"aliyundriver-webdav/m_container"
	"fyne.io/fyne/v2/app"
	"log"
	"os"
	"time"
)

func main() {
	lifecycle.InitBeforeConfig()
	lifecycle.InitConfig()
	args := os.Args
	for _, v := range args {
		println(v)
		if "-server" == v {

			for i := 0; i < 5; i++ {
				if !lifecycle.WebdavRunningStatus() {
					// 重试五次 可能因为网络问题没有成功
					bussiness.RunWebDav()
					if lifecycle.WebdavRunningStatus() {
						return
					}
					time.Sleep(time.Second * 1)
				}
			}
			return
		}
	}
	a := app.New()
	lifecycle.InitTheme(a)
	m_container.MApp = a
	go lifecycle.Process()

	//go systray.Run(m_systray.OpenSystray, nil)
	log.Println("程序已启动")
	a.Run()
}
