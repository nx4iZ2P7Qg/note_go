package filepath

import (
	"path/filepath"
	"testing"
)

func TestDemo(t *testing.T) {
	// 理想情况
	s := "C:/Windows/System32/drivers/etc/hosts"
	filepath.Dir(s)  // C:/Windows/System32/drivers/etc
	filepath.Base(s) // hosts

	// window 反斜线的解析
	s = "C:\\Windows\\System32\\drivers\\etc\\hosts"
	filepath.Dir(s)  // C:\Windows\System32\drivers\etc
	filepath.Base(s) // hosts

	// / 结尾，容易产生歧义的用法
	s = "C:/Windows/System32/drivers/etc/hosts/"
	filepath.Dir(s)  // C:/Windows/System32/drivers/etc/hosts
	filepath.Base(s) // hosts

	s = "C:/Windows/explorer.exe"
	filepath.Ext(s) // .exe

	filepath.IsAbs("c:/programs files") // true

	_, _ = filepath.Abs(".")

	s = "C:/Windows/System32/drivers/etc/hosts"
	_, _ = filepath.Split(s) // C:/Windows/System32/drivers/etc/	hosts

	//println(pr1)
	//println(pr2)
}
