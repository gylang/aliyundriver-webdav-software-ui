package ui

import (
	"aliyundriver-webdav/bussiness"
	m_container "aliyundriver-webdav/m_container"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

func OpenConfig() {
	form := make([]fyne.CanvasObject, 0)
	canEditFun := make([]func(), 0)
	notCanEditFun := make([]func(), 0)

	// 两个按钮
	var registerAutoStartBtn *widget.Button
	var removeAutoStartBtn *widget.Button
	registerAutoStartBtn = widget.NewButton("设置开机自启(需要管理员无需重复设置)", func() {
		// todo 启动状态判断
		err := bussiness.RegisterAutoStart()
		if err != nil {
			Error("设置开机自启失败", err.Error())
		} else {
			Msg("设置开机自启", "成功")
			registerAutoStartBtn.Disable()
			removeAutoStartBtn.Enable()
		}
	})

	removeAutoStartBtn = widget.NewButton("移除开机自启(需要管理员无需重复设置)", func() {
		// todo 启动状态判断
		err := bussiness.RemoveRegisterAutoStart()
		if err != nil {
			Error("移除开机自启失败", err.Error())
		} else {
			Msg("移除开机自启", "成功")
			registerAutoStartBtn.Enable()
			removeAutoStartBtn.Disable()
		}
	})
	autoStart := bussiness.QueryRegisterAutoStart()
	if autoStart {
		registerAutoStartBtn.Disable()
	} else {
		removeAutoStartBtn.Disable()
	}
	RunnStatus := widget.NewLabelWithData(m_container.MRunningStatus.StatusBinder)

	form = append(form, registerAutoStartBtn)
	form = append(form, removeAutoStartBtn)
	// AutoIndex
	form = append(form, canvas.NewText("自动生成主页(index,html)", color.Black))
	AutoIndexItem := widget.NewSelect([]string{"Y", "N"}, func(value string) {
		m_container.Config.AutoIndex = value
	})
	canEditFun = append(canEditFun, func() {
		AutoIndexItem.Enable()
	})
	notCanEditFun = append(notCanEditFun, func() {
		AutoIndexItem.Disable()
	})
	form = append(form, AutoIndexItem)

	//NoTrash
	form = append(form, canvas.NewText("永久删除文件而不是将其销毁", color.Black))
	NoTrashInput := widget.NewSelect([]string{"Y", "N"}, func(value string) {
		m_container.Config.NoTrash = value
	})
	canEditFun = append(canEditFun, func() {
		NoTrashInput.Enable()
	})
	notCanEditFun = append(notCanEditFun, func() {
		NoTrashInput.Disable()
	})
	form = append(form, NoTrashInput)

	//ReadOnly
	form = append(form, canvas.NewText("自读模式", color.Black))
	ReadOnlyInput := widget.NewSelect([]string{"Y", "N"}, func(value string) {
		m_container.Config.ReadOnly = value
	})
	canEditFun = append(canEditFun, func() {
		ReadOnlyInput.Enable()
	})
	notCanEditFun = append(notCanEditFun, func() {
		ReadOnlyInput.Disable()
	})
	form = append(form, ReadOnlyInput)

	//AuthUser
	form = append(form, canvas.NewText("webdav登录账号", color.Black))
	AuthUserInput := widget.NewEntry()
	canEditFun = append(canEditFun, func() {
		AuthUserInput.Enable()
	})
	notCanEditFun = append(notCanEditFun, func() {
		AuthUserInput.Disable()
	})
	AuthUserInput.Bind(binding.BindString(&(m_container.Config.AuthUser)))
	form = append(form, AuthUserInput)

	//AuthPassword
	form = append(form, canvas.NewText("webdav登录账号密码", color.Black))
	AuthPasswordInput := widget.NewEntry()
	canEditFun = append(canEditFun, func() {
		AuthPasswordInput.Enable()
	})
	notCanEditFun = append(notCanEditFun, func() {
		AuthPasswordInput.Disable()
	})
	AuthPasswordInput.Bind(binding.BindString(&(m_container.Config.AuthPassword)))
	form = append(form, AuthPasswordInput)

	//CacheSize
	form = append(form, canvas.NewText("目录条目缓存大小[默认值:1000]", color.Black))
	CacheSizeInput := widget.NewEntry()
	canEditFun = append(canEditFun, func() {
		CacheSizeInput.Enable()
	})
	notCanEditFun = append(notCanEditFun, func() {
		CacheSizeInput.Disable()
	})
	CacheSizeInput.Bind(binding.BindString(&(m_container.Config.CacheSize)))
	form = append(form, CacheSizeInput)

	//CacheTtl
	form = append(form, canvas.NewText("目录条目缓存过期时间(以秒为单位)[默认:600]", color.Black))
	CacheTtlInput := widget.NewEntry()
	canEditFun = append(canEditFun, func() {
		CacheTtlInput.Enable()
	})
	notCanEditFun = append(notCanEditFun, func() {
		CacheTtlInput.Disable()
	})
	CacheTtlInput.Bind(binding.BindString(&(m_container.Config.CacheTtl)))
	form = append(form, CacheTtlInput)

	//DomainId
	form = append(form, canvas.NewText(" Aliyun PDS domain id", color.Black))
	DomainIdInput := widget.NewEntry()
	canEditFun = append(canEditFun, func() {
		DomainIdInput.Enable()
	})
	notCanEditFun = append(notCanEditFun, func() {
		DomainIdInput.Disable()
	})
	DomainIdInput.Bind(binding.BindString(&(m_container.Config.DomainId)))
	form = append(form, DomainIdInput)

	//Host
	form = append(form, canvas.NewText("监听host[default: 0.0.0.0]", color.Black))
	HostInput := widget.NewEntry()
	canEditFun = append(canEditFun, func() {
		HostInput.Enable()
	})
	notCanEditFun = append(notCanEditFun, func() {
		HostInput.Disable()
	})
	HostInput.Bind(binding.BindString(&(m_container.Config.Host)))
	form = append(form, HostInput)

	//Port
	form = append(form, canvas.NewText("监听端口[default: 8080]", color.Black))
	PortInput := widget.NewEntry()
	canEditFun = append(canEditFun, func() {
		PortInput.Enable()
	})
	notCanEditFun = append(notCanEditFun, func() {
		PortInput.Disable()
	})
	PortInput.Bind(binding.BindString(&(m_container.Config.Port)))
	form = append(form, PortInput)

	//ReadBuffSize
	form = append(form, canvas.NewText("读取/下载缓冲区大小(字节)[默认值:10485760]  ", color.Black))
	ReadBuffSizeInput := widget.NewEntry()
	canEditFun = append(canEditFun, func() {
		ReadBuffSizeInput.Enable()
	})
	notCanEditFun = append(notCanEditFun, func() {
		ReadBuffSizeInput.Disable()
	})
	ReadBuffSizeInput.Bind(binding.BindString(&(m_container.Config.ReadBuffSize)))
	form = append(form, ReadBuffSizeInput)

	//RefreshToken
	form = append(form, canvas.NewText("阿里云盘refreshToken", &color.RGBA{R: 255, G: 0, B: 0, A: 255}))
	RefreshTokenInput := widget.NewEntry()
	canEditFun = append(canEditFun, func() {
		RefreshTokenInput.Enable()
	})
	notCanEditFun = append(notCanEditFun, func() {
		RefreshTokenInput.Disable()
	})
	RefreshTokenInput.Bind(binding.BindString(&(m_container.Config.RefreshToken)))
	form = append(form, RefreshTokenInput)

	//Root
	form = append(form, canvas.NewText("根目录路径[默认:/]", color.Black))
	RootInput := widget.NewEntry()
	canEditFun = append(canEditFun, func() {
		RootInput.Enable()
	})
	notCanEditFun = append(notCanEditFun, func() {
		RootInput.Disable()
	})
	RootInput.Bind(binding.BindString(&(m_container.Config.Root)))
	form = append(form, RootInput)

	// 两个按钮
	var startBtn *widget.Button
	var stopBtn *widget.Button
	startBtn = widget.NewButton("启动", func() {

		// 验证是否输入refresh_token
		if len(m_container.Config.RefreshToken) <= 0 {
			Error("参数错误", "RefreshToken不能为空")
			return
		}
		startBtn.Disable()
		stopBtn.Enable()
		Disable(notCanEditFun...)
		bussiness.RunWebDav()
		m_container.MRunningStatus.Running = true
		err := m_container.MRunningStatus.StatusBinder.Set("运行中")
		if err != nil {
			log.Fatal(err)
		}
	})

	stopBtn = widget.NewButton("停止(关闭失败可能是管理模式启动)", func() {
		startBtn.Enable()
		stopBtn.Disable()
		Enable(canEditFun...)
		err := bussiness.StopWebDav()
		if err != nil {
			Error("停止webdav服务", err.Error())
		} else {
			m_container.MRunningStatus.Running = false
			err = m_container.MRunningStatus.StatusBinder.Set("未启动")
			if err != nil {
				log.Fatal(err)
			}
		}

	})
	if m_container.MRunningStatus.Running {
		startBtn.Disable()
		Disable(notCanEditFun...)
	} else {
		stopBtn.Disable()
		Enable(canEditFun...)

	}
	startBtn.SetIcon(theme.ConfirmIcon())
	stopBtn.SetIcon(theme.CancelIcon())
	buttons := make([]*widget.Button, 0)
	buttons = append(buttons, startBtn)
	buttons = append(buttons, stopBtn)
	form = append(form, startBtn)
	form = append(form, stopBtn)
	grid := container.New(layout.NewGridLayout(2), form...)
	// 设置状态情况
	content := container.New(layout.NewVBoxLayout(), []fyne.CanvasObject{RunnStatus, grid}...)

	m_container.DefaultWindow().SetContent(content)
	m_container.DefaultWindow().CenterOnScreen()
	m_container.DefaultWindow().Resize(fyne.Size{Width: 400, Height: 600})
	m_container.DefaultWindow().Show()
}

func Disable(list ...func()) {
	for _, v := range list {
		v()
	}
}
func Enable(list ...func()) {
	for _, v := range list {
		v()
	}
}
