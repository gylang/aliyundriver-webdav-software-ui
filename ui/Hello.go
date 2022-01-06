package ui

import (
	"aliyundriver-webdav/m_container"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func OpenHello() {

	m_container.DefaultWindow().Resize(fyne.Size{Height: 50, Width: 50})
	m_container.DefaultWindow().SetContent(widget.NewLabel("Hello World!"))
	m_container.DefaultWindow().Show()
}
