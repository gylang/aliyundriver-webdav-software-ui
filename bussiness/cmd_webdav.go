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
	"syscall"
)

func RunWebDav() {

	params := parseParams(*m_container.Config)
	command := exec.Command(constant.WebdavPath(), params...)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	log.Println(command.String())
	logFile, err := os.OpenFile(constant.WebdavLogsPath(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
	if err != nil {
		log.Println("记录webdav执行日志文件打开失败 " + err.Error())
	} else {
		command.Stderr = logFile
		command.Stdout = logFile
	}
	command.Start()

}

func StopWebDav() {

	if runtime.GOOS == "windows" {
		if runtime.GOOS == "windows" {
			cmd := exec.Command("taskkill.exe", "/F", "/T", "/IM", constant.WinAliyunDriveWebdav)
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			log.Println(cmd.String())
			res, err := cmd.CombinedOutput()
			result := util.ConvertToString(string(res), "gbk", "utf8")
			log.Println(result)
			if err != nil {
				log.Println(err.Error())
			}
		}

	}
}

func parseParams(config domain.AliWebDavConfig) []string {

	cmd := make([]string, 0)
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
