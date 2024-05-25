package main

import (
	"github.com/buhuang28/mini-tool/logs"
	"github.com/buhuang28/mini-tool/ui"
	"github.com/ying32/govcl/vcl"
)

func init() {
	logs.LogInit()
}

//var (
//	gpToken winapi.ULONG_PTR
//)

func main() {
	//_, err := gdiplus.Startup(&gpToken, nil, nil)
	//if err != nil {
	//	if err != nil {
	//		vcl.ShowMessage("gdi+ init error:" + err.Error())
	//	}
	//} else {
	//	fmt.Println("gdi+ init success")
	//	defer gdiplus.Shutdown(gpToken)
	//}
	vcl.Application.Initialize()
	vcl.Application.CreateForm(&ui.MainForm)
	vcl.Application.SetMainFormOnTaskBar(false)
	vcl.Application.Run()
}
