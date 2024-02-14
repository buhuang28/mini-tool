package per

import (
	"fmt"
	"github.com/buhuang28/mini-tool/cst"
	"github.com/shirou/gopsutil/v4/cpu"
	log "github.com/sirupsen/logrus"
	"time"
)

func GetCpuUsedRate() string {
	percent, err := cpu.Percent(time.Second*1, false)
	if err != nil {
		log.Error(err)
		return cst.ERROR
	}

	var sum float64
	for _, v := range percent {
		sum += v
	}

	cpuRate := sum / float64(len(percent))
	return fmt.Sprintf("%.2f", cpuRate)
}
