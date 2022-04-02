package web

import (
	"aliyundriver-webdav/bussiness"
	"aliyundriver-webdav/constant"
	"aliyundriver-webdav/domain"
	"aliyundriver-webdav/lifecycle"
	"aliyundriver-webdav/m_container"
	"aliyundriver-webdav/util"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"
)

var srv *http.Server
var closeChan chan bool

func OpenWebServer() {

	if m_container.Config.WebPost <= 0 {
		m_container.Config.WebPost = 35640
	}
	if JustOpenPage() {
		return
	}
	// 关闭旧进程
	lifecycle.CloseOld()
	http.HandleFunc("/console", console)
	http.HandleFunc("/css/", css)
	http.HandleFunc("/js/", js)
	http.HandleFunc("/font/", font)
	http.HandleFunc("/status", status)
	http.HandleFunc("/start", start)
	http.HandleFunc("/stop", stop)
	http.HandleFunc("/closeApp", closeApp)
	http.HandleFunc("/registerAutoStart", registerAutoStart)
	http.HandleFunc("/removeRegisterAutoStart", removeRegisterAutoStart)

	closeChan = make(chan bool)
	go func() {
		log.Println("启动web服务")
		srv = &http.Server{Addr: fmt.Sprintf("127.0.0.1:%d", m_container.Config.WebPost)}
		srv.ListenAndServe()
	}()
	OpenWebPage()
	select {
	case <-closeChan:
		fmt.Println("接听到结束时间")
		timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer timeoutCancel()
		srv.Shutdown(timeoutCtx)
		return
	}
}

func closeApp(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "关闭程序")
	lifecycle.CloseOld()
	closeChan <- true
}
func JustOpenPage() bool {
	// 判断是否已经启动服务
	if len(bussiness.ListProcess("IMAGENAME eq "+constant.AppProcessName)) > 0 {
		// 判断旧的进程是否成功启动web服务
		resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/status", m_container.Config.WebPost))
		if err != nil {
			log.Println("旧服务响应结果: " + err.Error())
		} else if resp.StatusCode == 200 {
			log.Println("旧进程响应200 成功响应")
			return true
		}
	}
	return false
}

func OpenWebPage() {
	command := exec.Command("cmd", "/c", "start", fmt.Sprintf("http://127.0.0.1:%d/console", m_container.Config.WebPost))
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	command.Start()
}

func removeRegisterAutoStart(w http.ResponseWriter, r *http.Request) {

	err := bussiness.RemoveRegisterAutoStart()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
	} else {
		w.WriteHeader(200)
		fmt.Fprintf(w, "取消开机自启成功")
	}
}

func registerAutoStart(w http.ResponseWriter, r *http.Request) {
	err := bussiness.RegisterAutoStart()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
	} else {
		w.WriteHeader(200)
		fmt.Fprintf(w, "注册开机自启成功")
	}
}

func stop(w http.ResponseWriter, r *http.Request) {

	err := bussiness.StopWebDav()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
	} else {
		w.WriteHeader(200)
		fmt.Fprintf(w, "停止服务成功")
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	var p domain.AliWebDavConfig
	bytes, err := io.ReadAll(r.Body)
	if err == nil {
		json.Unmarshal(bytes, &p)
		m_container.Config.AutoIndex = p.AutoIndex
		m_container.Config.NoTrash = p.NoTrash
		m_container.Config.ReadOnly = p.ReadOnly
		m_container.Config.Version = p.Version
		m_container.Config.AuthPassword = p.AuthPassword
		m_container.Config.AuthUser = p.AuthUser
		m_container.Config.CacheSize = p.CacheSize
		m_container.Config.CacheTtl = p.CacheTtl
		m_container.Config.DomainId = p.DomainId
		m_container.Config.Host = p.Host
		m_container.Config.Port = p.Port
		m_container.Config.ReadBuffSize = p.ReadBuffSize
		m_container.Config.SyncRefreshToken = p.SyncRefreshToken
		m_container.Config.RefreshToken = p.RefreshToken
		m_container.Config.Root = p.Root
		bussiness.RunWebDav()
		w.WriteHeader(200)
		fmt.Fprintf(w, "尝试启动成功")
	} else {
		w.WriteHeader(500)
		fmt.Fprintf(w, err.Error())
	}

}

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, strconv.FormatBool(m_container.MRunningStatus.Running))
}

func console(w http.ResponseWriter, r *http.Request) {

	path := util.OsPathConverter(util.CurrentDirectory() + "/template/console.html")
	t1, err := template.ParseFiles(path)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, err.Error())
	} else {
		t1.Execute(w, m_container.Config)
	}

}
func css(w http.ResponseWriter, r *http.Request) {

	path := util.OsPathConverter(util.CurrentDirectory() + "/template" + r.URL.Path)
	w.Header().Add("content-type", "text/css")
	WriteStaticFile(w, path)
}

func js(w http.ResponseWriter, r *http.Request) {
	path := util.OsPathConverter(util.CurrentDirectory() + "/template" + r.URL.Path)
	w.Header().Add("content-type", "application/javascript")
	WriteStaticFile(w, path)
}

func font(w http.ResponseWriter, r *http.Request) {
	path := util.OsPathConverter(util.CurrentDirectory() + "/template" + r.URL.Path)
	w.Header().Add("content-type", "text/html")
	WriteStaticFile(w, path)
}
func WriteStaticFile(w http.ResponseWriter, path string) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0444)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, err.Error())
	} else {
		w.WriteHeader(200)
		bytes, err := io.ReadAll(file)
		if err != nil {
			w.WriteHeader(404)
			fmt.Fprint(w, err.Error())
		} else {
			w.Write(bytes)
		}
	}
}
