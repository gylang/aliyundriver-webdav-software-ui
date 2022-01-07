package m_container

import "fyne.io/fyne/v2/data/binding"

type RunningStatus struct {
	Running      bool
	StatusBinder binding.String
}

var MRunningStatus = RunningStatus{Running: false, StatusBinder: binding.NewString()}

var StartChan chan bool
