package per

import (
	"fmt"
	"github.com/StackExchange/wmi"
	"github.com/buhuang28/mini-tool/cst"
	log "github.com/sirupsen/logrus"
)

type Win32_PerfFormattedData_Tcpip_NetworkInterface struct {
	Name                string
	BytesReceivedPerSec uint32
	BytesSentPerSec     uint32
}

type NetworkSpeed struct {
	RecvSpeed string
	SendSpeed string
}

var (
	dst []Win32_PerfFormattedData_Tcpip_NetworkInterface
	q   = wmi.CreateQuery(&dst, `` /*`WHERE Name = "Realtek PCIe GBE Family Controller"`*/)
)

func GetNetworkSpeed() NetworkSpeed {
	var d []Win32_PerfFormattedData_Tcpip_NetworkInterface
	err := wmi.Query(q, &d)
	if err != nil {
		log.Error(err)
		return NetworkSpeed{
			RecvSpeed: cst.ERROR,
			SendSpeed: cst.ERROR,
		}
	}

	var recv, send uint32

	for _, v := range d {
		recv += v.BytesReceivedPerSec
		send += v.BytesSentPerSec
	}

	recvKbs := recv / 1024
	sendKbs := send / 1024

	recvSpeed, sendSpeed := "", ""

	if recvKbs > 1024 {
		recvSpeed = fmt.Sprintf("%.2fMb/s", recvKbs/1024)
	} else {
		recvSpeed = fmt.Sprintf("%dKb/s", recvKbs)
	}

	if sendKbs > 1024 {
		sendSpeed = fmt.Sprintf("%.2fMb/s", sendKbs/1024)
	} else {
		sendSpeed = fmt.Sprintf("%dKb/s", sendKbs)
	}

	return NetworkSpeed{
		RecvSpeed: recvSpeed,
		SendSpeed: sendSpeed,
	}
}
