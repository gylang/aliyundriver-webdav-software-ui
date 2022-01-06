package bussiness

import (
	"aliyundriver-webdav/constant"
	"aliyundriver-webdav/domain"
	"aliyundriver-webdav/m_container"
	"aliyundriver-webdav/util"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func RunWebDav() {

	params := parseParams(*m_container.Config)
	command := exec.Command("cmd.exe", append([]string{"/c", util.CurrentDirectory() + string(os.PathSeparator) + "bin" + string(os.PathSeparator) + "aliyundriver_start.vbs"}, params...)...)
	logFile, err := os.OpenFile(constant.WebdavLogsPath(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
	if err != nil {
		log.Println("记录webdav执行日志文件打开失败 " + err.Error())
	} else {
		command.Stderr = logFile
		command.Stdout = logFile
	}
	log.Println(command.String())
	command.Start()

}

func StopWebDav() {

	if runtime.GOOS == "windows" {
		if runtime.GOOS == "windows" {
			cmd := exec.Command("taskkill.exe", "/F", "/T", "/IM", constant.WinAliyunDriveWebdav)
			log.Println(cmd.String())
			_, err := cmd.CombinedOutput()
			if err != nil {
				log.Println(err.Error())
			}

		}

	}
}

func parseParams(config domain.AliWebDavConfig) []string {

	cmd := make([]string, 0)
	cmd = append(cmd, util.CurrentDirectory()+string(os.PathSeparator)+"bin"+string(os.PathSeparator)+"aliyundrive-webdav.exe")
	if "Y" == config.AutoIndex {
		cmd = append(cmd, "-I")
	}
	if "Y" == config.NoTrash {
		cmd = append(cmd, "--no-trash")
	}
	if "Y" == config.ReadOnly {
		cmd = append(cmd, "--read-only")
	}
	if len(config.AuthPassword) > 0 {
		cmd = append(cmd, "-W", config.AuthPassword)
	}
	if len(config.AuthUser) > 0 {
		cmd = append(cmd, "-U", config.AuthUser)
	}
	if len(config.CacheSize) > 0 {
		cmd = append(cmd, "--cache-size", config.CacheSize)
	}
	if len(config.CacheTtl) > 0 {
		cmd = append(cmd, "--cache-ttl", config.CacheTtl)
	}
	if len(config.DomainId) > 0 {
		cmd = append(cmd, "--domain-id", config.DomainId)
	}
	if len(config.Host) > 0 {
		cmd = append(cmd, "--host", config.Host)
	}
	if len(config.Port) > 0 {
		cmd = append(cmd, "-p", config.Port)
	}
	if len(config.ReadBuffSize) > 0 {
		cmd = append(cmd, "-S", config.ReadBuffSize)
	}
	if len(config.RefreshToken) > 0 {
		cmd = append(cmd, "-r", config.RefreshToken)
	}
	if len(config.Root) > 0 {
		cmd = append(cmd, "--root", config.Root)
	}
	if len(config.WorkDir) > 0 {
		cmd = append(cmd, "-w", config.WorkDir)
	}
	return cmd

}
