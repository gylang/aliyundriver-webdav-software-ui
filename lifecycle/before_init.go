package lifecycle

import (
	"aliyundriver-webdav/bussiness"
	"aliyundriver-webdav/constant"
	"aliyundriver-webdav/domain"
	"aliyundriver-webdav/m_container"
	"aliyundriver-webdav/util"
	"encoding/json"
	"github.com/gookit/goutil/sysutil/process"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func InitConfig() {

}
func InitBeforeConfig() {
	//closeOld()
	initLog()
	initFont()
	intConfig()
	MonitorWebdavStatus()
}

func closeOld() {
	listProcess := bussiness.ListProcess("IMAGENAME eq aliyundriver-webdav.exe")
	for _, v := range listProcess {
		if v.Pid != process.PID() {
			if v.Pid > 0 {
				bussiness.KillProcessByPid(strconv.Itoa(v.Pid))
			}
		}
	}
}

func intConfig() {
	config := domain.AliWebDavConfig{}
	m_container.Config = &config
	confPath := constant.WebdavConfigPath()
	file, _ := os.OpenFile(confPath, os.O_RDONLY, 0777)
	bytes, _ := io.ReadAll(file)
	json.Unmarshal(bytes, &config)
	file.Close()
	go func() {
		for {
			<-time.NewTicker(time.Second * 30).C
			jsonConf, err := json.Marshal(&config)
			if err == nil {
				file, err = os.OpenFile(confPath, os.O_RDWR|os.O_TRUNC, 0777)
				if err != nil {
					continue
				}
				_, err := file.Write(jsonConf)
				if err != nil {
					log.Println(err.Error())
				}
			}
		}
	}()
}

func initFont() {
	// 加载字体
	log.Println(constant.FontPath())
	os.Setenv("FYNE_FONT", constant.FontPath())
}

func initLog() {
	LogPath := constant.LogPath()
	if _, err := os.Stat(LogPath); os.IsNotExist(err) {
		util.MakeDir(LogPath)
	}
	file := LogPath + string(os.PathSeparator) + time.Now().Format("2006-01-02") + ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[GiLang-ali-webdav]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("---日志加载完成---")
	log.Print("---日志加载完成---")
	return
}
