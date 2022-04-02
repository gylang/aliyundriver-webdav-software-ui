package bussiness

import (
	"aliyundriver-webdav/constant"
	"aliyundriver-webdav/m_container"
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

var Writing = sync.Mutex{}

func WriteConfig() {

	Writing.Lock()
	// 加锁成功
	log.Println("写入加锁成功")
	defer Writing.Unlock()
	jsonConf, err := json.Marshal(&m_container.Config)
	if err == nil {
		file, err := os.OpenFile(constant.WebdavConfigPath(), os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
		if err != nil {
			log.Println(err.Error())
			return
		}
		_, err = file.Write(jsonConf)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func TimerPersistence() {

	// 启动写入一次
	WriteConfig()
	go func() {
		for {
			<-time.NewTicker(time.Second * 30).C
			WriteConfig()
		}
	}()
}
