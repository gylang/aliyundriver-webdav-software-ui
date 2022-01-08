package bussiness

import (
	"aliyundriver-webdav/domain"
	"github.com/gookit/goutil/strutil"
	"log"
	"os/exec"
	"syscall"
)

func ListProcess(text string) []domain.MProcess {
	cmd := exec.Command("cmd", "/c", "tasklist", "/FI", text)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return Parse2MProcess(string(out))
}

func ListAllProcess() []domain.MProcess {
	cmd := exec.Command("cmd", "/c", "tasklist")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return Parse2MProcess(string(out))
}

func killProcess(text string) {
	cmd := exec.Command("cmd", "/c", "taskkill", "/F", "/T", "/FI", text)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))
}

func KillProcessByPid(text string) {
	cmd := exec.Command("taskkill", "/F", "/T", "/PID", text)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))
}

func Parse2MProcess(s string) []domain.MProcess {
	mProcesses := make([]domain.MProcess, 0)
	lines := strutil.Split(s, "\n")
	for _, v := range lines {
		items := strutil.Split(v, " ")
		if len(items) > 5 {
			mProcesses = append(mProcesses, domain.MProcess{Name: items[0], Pid: strutil.MustInt(items[1]),
				Type: items[2], SessionNum: strutil.MustInt(items[3]), Memory: strutil.MustInt(items[4]), MemoryUnit: items[5]})
		}

	}
	return mProcesses
}
