package ui

import (
	"fmt"
	"github.com/buhuang28/mini-tool/per"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"time"
)

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("mini-tool")
	f.SetWidth(300)
	f.SetHeight(300)
	f.SetBorderStyle(types.BsSingle)
	//f.SetBorderStyle(types.BsNone)
	//f.SetAlphaBlend(true)
	//f.SetAlphaBlendValue(200)
	f.ScreenCenter()
	f.SetAutoSize(false)
	f.SetDoubleBuffered(true)
	f.SetColor(colors.ClAzure)

	f.UIPage = vcl.NewPageControl(f)
	f.UIPage.SetParent(f)
	f.UIPage.SetWidth(f.Width())
	f.UIPage.SetHeight(f.Height())
	f.UIPage.SetLeft(0)
	f.UIPage.SetTop(0)

	f.Performance = vcl.NewTabSheet(f)
	f.Performance.SetParent(f.UIPage)
	f.Performance.SetCaption("性能监控")

	f.CpuText = f.InitLabel("CPU使用率:", 10, 10)
	f.Cpu = f.InitLabelOnRight(f.CpuText, "56%")
	f.Cpu.Font().SetColor(colors.ClGreen)
	f.MemText = f.InitLabelOnFormRight(f.CpuText, "内存使用率:")
	f.Mem = f.InitLabelOnRight(f.MemText, "56%")
	f.Mem.Font().SetColor(colors.ClGreen)
	f.NetRecvText = f.InitLabelOnButtom(f.CpuText, "网 络 下 行:")
	f.NetRecv = f.InitLabelOnRight(f.NetRecvText, "1024Kb/s")
	f.NetRecv.Font().SetColor(colors.ClGreen)
	f.NetSendText = f.InitLabelOnButtom(f.MemText, "网 络 上 行:")
	f.NetSend = f.InitLabelOnRight(f.NetSendText, "1024Kb/s")
	f.NetSend.Font().SetColor(colors.ClGreen)
	f.DiskReadText = f.InitLabelOnButtom(f.NetRecvText, "硬 盘 读 取:")
	f.DiskRead = f.InitLabelOnRight(f.DiskReadText, "1024Kb/s")
	f.DiskRead.Font().SetColor(colors.ClGreen)

	f.DiskWriteText = f.InitLabelOnButtom(f.NetSendText, "硬 盘 写 入:")
	f.DiskWrite = f.InitLabelOnRight(f.DiskWriteText, "1024Kb/s")
	f.DiskWrite.Font().SetColor(colors.ClGreen)

	f.Process = vcl.NewTabSheet(f)
	f.Process.SetParent(f.UIPage)
	f.Process.SetCaption("进程列表")

	f.ProcessList = vcl.NewListView(f)
	f.ProcessList.SetParent(f.Process)
	f.ProcessList.SetViewStyle(types.VsReport)
	f.ProcessList.SetBorderStyle(types.BsSingle)
	f.ProcessList.SetReadOnly(true)
	f.ProcessList.SetRowSelect(true)
	//f.WinList.SetMultiSelect(false)
	f.ProcessList.SetColor(colors.ClAzure)
	f.ProcessList.SetAlign(types.AlNone)
	f.ProcessList.SetWidth(300)
	f.ProcessList.SetHeight(250)

	addColV2(f.ProcessList, []Item{NewItem("PID", 60), NewItem("进程名", 240)})

	f.Kill = vcl.NewButton(f)
	f.Kill.SetParent(f.Process)
	f.Kill.SetTop(f.ProcessList.Top() + f.ProcessList.Height())
	f.Kill.SetLeft(f.ProcessList.Left())
	f.Kill.SetWidth(f.UIPage.Width()/2 - 5)
	f.Kill.SetCaption("终止该程序启动")

	f.LimitNetwork = vcl.NewButton(f)
	f.LimitNetwork.SetParent(f.Process)
	f.LimitNetwork.SetTop(f.ProcessList.Top() + f.ProcessList.Height())
	f.LimitNetwork.SetLeft(f.Kill.Left() + f.Kill.Width() + 1)
	f.LimitNetwork.SetWidth(f.UIPage.Width()/2 - 5)
	f.LimitNetwork.SetCaption("禁止该程序联网")

	go func() {
		for {
			select {
			case <-time.After(time.Second * 3):
				process := per.GetProcess()
				vcl.ThreadSync(func() {
					for _, v := range process {
						item := f.ProcessList.Items().Add()
						item.SubItems().Add(fmt.Sprintf("%d", v.Pid))
						item.SubItems().Add(v.Name)
					}
				})
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(time.Second * 1):
				cpuUsedRate := per.GetCpuUsedRate() + "%"
				memUsedRate := per.GetMemUsedRate() + "%"
				netSpeed := per.GetNetworkSpeed()
				vcl.ThreadSync(func() {
					f.Cpu.SetCaption(cpuUsedRate)
					f.Mem.SetCaption(memUsedRate)
					f.NetRecv.SetCaption(netSpeed.RecvSpeed)
					f.NetSend.SetCaption(netSpeed.SendSpeed)
				})
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(time.Second * 1):
				netSpeed := per.GetNetworkSpeed()
				vcl.ThreadSync(func() {
					f.NetRecv.SetCaption(netSpeed.RecvSpeed)
					f.NetSend.SetCaption(netSpeed.SendSpeed)
				})
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(time.Second * 1):
				diskSpeed := per.GetDiskWRSpeed()
				vcl.ThreadSync(func() {
					f.DiskRead.SetCaption(diskSpeed.ReadSpeed)
					f.DiskWrite.SetCaption(diskSpeed.WriteSpeed)
				})
			}
		}
	}()

}

func (f *TMainForm) InitLabel(text string, top, left int32) *vcl.TLabel {
	ctl := vcl.NewLabel(f)
	ctl.SetParent(f.Performance)
	ctl.SetCaption(text)
	ctl.SetTop(top)
	ctl.SetLeft(left)
	ctl.SetAutoSize(true)
	return ctl
}

func (f *TMainForm) InitLabelOnRight(one *vcl.TLabel, text string) *vcl.TLabel {
	ctl := vcl.NewLabel(f)
	ctl.SetParent(f.Performance)
	ctl.SetCaption(text)
	ctl.SetTop(one.Top())
	ctl.SetLeft(one.Left() + one.Width() + 5)
	ctl.SetAutoSize(true)
	return ctl
}

func (f *TMainForm) InitLabelOnFormRight(one *vcl.TLabel, text string) *vcl.TLabel {
	ctl := vcl.NewLabel(f)
	ctl.SetParent(f.Performance)
	ctl.SetCaption(text)
	ctl.SetTop(one.Top())
	ctl.SetLeft(f.Width()/2 + 1)
	ctl.SetAutoSize(true)
	return ctl
}

func (f *TMainForm) InitLabelOnButtom(one *vcl.TLabel, text string) *vcl.TLabel {
	ctl := vcl.NewLabel(f)
	ctl.SetParent(f.Performance)
	ctl.SetCaption(text)
	ctl.SetTop(one.Top() + 25)
	ctl.SetLeft(one.Left())
	ctl.SetAutoSize(true)
	ctl.SetWidth(one.Width())
	return ctl
}
