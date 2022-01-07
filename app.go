package main

import (
	"aliyundriver-webdav/bussiness"
	"aliyundriver-webdav/lifecycle"
	"aliyundriver-webdav/m_container"
	"fyne.io/fyne/v2/app"
	"log"
	"os"
)

func main() {
	lifecycle.InitBeforeConfig()
	lifecycle.InitConfig()
	args := os.Args
	for _, v := range args {
		println(v)
		if "-server" == v {
			bussiness.RunWebDav()
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
