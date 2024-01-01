package ui

import "github.com/ying32/govcl/vcl"

var MainForm *TMainForm

type TMainForm struct {
	*vcl.TForm

	UIPage *vcl.TPageControl

	Performance   *vcl.TTabSheet
	CpuText       *vcl.TLabel
	Cpu           *vcl.TLabel
	MemText       *vcl.TLabel
	Mem           *vcl.TLabel
	DiskReadText  *vcl.TLabel
	DiskWriteText *vcl.TLabel
	DiskRead      *vcl.TLabel
	DiskWrite     *vcl.TLabel
	NetRecvText   *vcl.TLabel
	NetSendText   *vcl.TLabel
	NetRecv       *vcl.TLabel
	NetSend       *vcl.TLabel

	Process      *vcl.TTabSheet
	ProcessList  *vcl.TListView
	LimitNetwork *vcl.TButton
	Kill         *vcl.TButton
}
