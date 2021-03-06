package main

import (
	"aliyundriver-webdav/bussiness"
	"aliyundriver-webdav/lifecycle"
	"aliyundriver-webdav/m_container"
	"aliyundriver-webdav/web"
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
	if m_container.Config.WebAccess {
		web.OpenWebServer()
	} else {
		func() {
			defer func() {
				// 启动失败异常处理
				if p := recover(); p != nil {
					// ui打不开 使用页面方式
					log.Println(p)
					log.Println("启动失败, 打开web页面")
					web.OpenWebServer()
				}
			}()
			a := app.New()
			lifecycle.InitTheme(a)
			m_container.MApp = a
			go lifecycle.Process()
			a.Run()

		}()
	}
}
