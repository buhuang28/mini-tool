package main

import (
	"flag"
	"fmt"
	"github.com/buhuang28/mini-tool/caches"
	"github.com/buhuang28/mini-tool/logs"
	"github.com/buhuang28/mini-tool/per"
	"github.com/buhuang28/mini-tool/ui"
	"github.com/tryor/gdiplus"
	"github.com/tryor/winapi"
	"github.com/ying32/govcl/vcl"
	"time"
)

func init() {
	logs.LogInit()
	go func() {
		for {
			select {
			case <-time.After(time.Second * 3):
				per.KillProcess()
			}
		}
	}()
}

var cfgPath = flag.String("f", "./config.json", "the config file")

var (
	gpToken winapi.ULONG_PTR
)

func main() {
	flag.Parse()

	caches.ConfigPath = *cfgPath
	_, err := gdiplus.Startup(&gpToken, nil, nil)
	if err != nil {
		if err != nil {
			vcl.ShowMessage("gdi+ init error:" + err.Error())
		}
	} else {
		fmt.Println("gdi+ init success")
		defer gdiplus.Shutdown(gpToken)
	}
	vcl.Application.Initialize()
	vcl.Application.CreateForm(&ui.MainForm)
	vcl.Application.SetMainFormOnTaskBar(false)
	vcl.Application.Run()
}
