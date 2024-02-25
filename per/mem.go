package per

import (
	"fmt"
	"github.com/buhuang28/mini-tool/cst"
	"github.com/shirou/gopsutil/v4/mem"
	log "github.com/sirupsen/logrus"
)

func GetMemUsedRate() string {
	m, err := mem.VirtualMemory()
	if err != nil {
		log.Error(err)
		return cst.ERROR
	}
	return fmt.Sprintf("%.2f", m.UsedPercent)
}
