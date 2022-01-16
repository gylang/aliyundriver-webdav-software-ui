package constant

import (
	"aliyundriver-webdav/util"
	"os"
	"time"
)

const WinAliyunDriveWebdav = "aliyundrive-webdav.exe"

func WebdavLogsPath() string {
	return util.CurrentDirectory() + string(os.PathSeparator) + "logs" + string(os.PathSeparator) + time.Now().Format("2006-01-02") + "-webdav-executor.txt"
}
func WebdavConfigPath() string {
	return util.CurrentDirectory() + string(os.PathSeparator) + "bin" + string(os.PathSeparator) + "webdav-conf.json"
}
func FontPath() string {
	return util.CurrentDirectory() + string(os.PathSeparator) + "resource" + string(os.PathSeparator) + "simhei.ttf"
}
func LogPath() string {
	return util.CurrentDirectory() + string(os.PathSeparator) + "logs"
}
func IcoPath() string {
	return util.CurrentDirectory() + string(os.PathSeparator) + "bin" + string(os.PathSeparator) + "bitbug_favicon.ico"
}
func IconJpgPath() string {
	return util.CurrentDirectory() + string(os.PathSeparator) + "bin" + string(os.PathSeparator) + "icon.jpg"
}
func AppServerModeCmd() string {
	return util.CurrentDirectory() + string(os.PathSeparator) + "aliyundriver-webdav.exe -server"
}
func WebdavPath() string {
	return util.CurrentDirectory() + string(os.PathSeparator) + "bin" + string(os.PathSeparator) + "aliyundrive-webdav.exe"
}
func WebdavRefreshTokenPath() string {
	return util.CurrentDirectory() + string(os.PathSeparator) + "bin" + string(os.PathSeparator) + "refresh_token"
}
