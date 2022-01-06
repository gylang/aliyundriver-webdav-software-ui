package constant

import (
	"aliyundriver-webdav/util"
	"os"
	"time"
)

const WinAliyunDriveWebdav = "aliyundrive-webdav.exe"

func WebdavLogsPath() string {
	dir := util.CurrentDirectory()
	return dir + string(os.PathSeparator) + "logs" + string(os.PathSeparator) + time.Now().Format("2006-01-02150405") + "-webdav-executor.txt"
}
