package bussiness

import (
	"log"
	"os"
	"os/exec"
)

func RegisterAutoStart() {
	dir, _ := os.Getwd()
	// "reg add HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /v AUTORUN /t REG_SZ /d C:\autorun.exe /f
	cmd := exec.Command("REG", "ADD", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", "/v", "go_auto_start", "/t", "REG_SZ", "/d", dir+string(os.PathSeparator)+"aliyundriver-webdav.exe -server", "/f")
	log.Println(cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(out))
}

func RemoveRegisterAutoStart() {

	// "reg add HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /v AUTORUN /t REG_SZ /d C:\autorun.exe /f
	cmd := exec.Command("REG", "DELETE", "HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", "/v", "go_auto_start")
	log.Println(cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(out))
}
