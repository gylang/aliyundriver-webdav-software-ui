package bussiness

import (
	"aliyundriver-webdav/constant"
	"aliyundriver-webdav/domain"
	"github.com/gookit/goutil/sysutil/process"
	"reflect"
	"testing"
)

func Test_listAllProcess(t *testing.T) {
	tests := []struct {
		name string
		want []domain.MProcess
	}{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListAllProcess(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listAllProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listProcess(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []domain.MProcess
	}{
		{name: constant.AppProcessName, args: args{text: "IMAGENAME eq aliyundriver-webdav.exe"}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListProcess(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("进程id = %d", process.PID())
				t.Errorf("listProcess() = %v, want %v", got, tt.want)
			}
		})
	}
}
