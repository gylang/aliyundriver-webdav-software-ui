package ui

import (
	"aliyundriver-webdav/m_container"
	"fmt"
	"fyne.io/fyne/v2/dialog"
	"log"
)

func Error(title string, msg string) {

	dialog.NewError(fmt.Errorf(title+": "+msg), m_container.DefaultWindow()).Show()
}
func Msg(title string, msg string) {

	dialog.NewConfirm(title, msg, func(b bool) {

	}, m_container.DefaultWindow()).Show()
}

func CloseError(title string, msg string) {
	log.Println(msg)
	dialog.NewConfirm(title, msg, func(b bool) {
		m_container.DefaultWindow().Close()
	}, m_container.DefaultWindow()).Show()
}
