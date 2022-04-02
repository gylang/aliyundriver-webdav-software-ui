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
	bussiness.TimerPersistence()
}

func CloseOld() {
	listProcess := bussiness.ListProcess("IMAGENAME eq " + constant.AppProcessName)
	for _, v := range listProcess {
		if v.Pid != process.PID() {
			if v.Pid > 0 {
				bussiness.KillProcessByPid(strconv.Itoa(v.Pid))
			}
		}
	}
}

func intConfig() {
	config := domain.AliWebDavConfig{WorkDir: constant.WebdavRefreshTokenPathFolder()}
	m_container.Config = &config
	confPath := constant.WebdavConfigPath()
	file, _ := os.OpenFile(confPath, os.O_RDONLY, 0777)
	bytes, err := io.ReadAll(file)
	if err == nil {
		json.Unmarshal(bytes, &config)
	}
	refreshToken, err := os.OpenFile(constant.WebdavRefreshTokenPath(), os.O_RDONLY, 0444)
	if err != nil {
		log.Println("打开refresh_token 失败")
	} else {
		rf, _ := io.ReadAll(refreshToken)
		if rf != nil && string(rf) != m_container.Config.SyncRefreshToken {
			rfStr := string(rf)
			m_container.Config.SyncRefreshToken = rfStr
			m_container.Config.RefreshToken = rfStr
		}
		refreshToken.Close()
	}

	file.Close()

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
