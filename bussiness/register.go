package bussiness

import (
	"aliyundriver-webdav/constant"
	"aliyundriver-webdav/util"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func RegisterAutoStart() error {
	// reg ADD HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /v "D:\software\aliyundriver-webdav\aliyundriver-webdav.exe -server" /f
	cmd := exec.Command("REG", "ADD", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", "/v", "go_auto_start", "/t", "REG_SZ", "/d", constant.AppServerModeCmd(), "/f")
	log.Println(cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Println(err.Error())
		return err
	}
	result := util.ConvertToString(string(out), "gbk", "utf8")
	if strings.Contains(result, "操作成功完成") {

		return nil
	}
	log.Println("执行命令结果 ： " + result)

	return fmt.Errorf(result)
}

func RemoveRegisterAutoStart() error {

	// REG DELETE HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /v go_auto_start
	cmd := exec.Command("REG", "DELETE", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", "/v", "go_auto_start", "/f")
	log.Println(cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	result := util.ConvertToString(string(out), "gbk", "utf8")
	log.Println("执行命令结果 ： " + result)
	if strings.Contains(result, "操作成功完成") {
		return nil
	}
	return fmt.Errorf(result)
}
func QueryRegisterAutoStart() bool {

	// REG QUERY HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /v go_auto_start
	cmd := exec.Command("REG", "QUERY", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", "/v", "go_auto_start")
	log.Println(cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	result := util.ConvertToString(string(out), "gbk", "utf8")
	log.Println("执行命令结果 ： " + result)
	if strings.Contains(result, constant.AppServerModeCmd()) {
		return true
	}
	return false
}
