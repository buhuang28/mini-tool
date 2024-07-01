package per

import (
	"fmt"
	"github.com/buhuang28/mini-tool/cst"
	"github.com/shirou/gopsutil/v4/disk"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	LastReadBytes  uint64
	LastWriteBytes uint64
	LastGetTime    int64
)

type DiskRWSpeed struct {
	WriteSpeed string
	ReadSpeed  string
}

func GetDiskWRSpeed() DiskRWSpeed {
	if (LastReadBytes == 0 && LastWriteBytes == 0) || (time.Now().Unix()-LastGetTime > 1) {
		counters, err := disk.IOCounters()
		if err != nil {
			log.Error(err)
			return DiskRWSpeed{
				WriteSpeed: cst.ERROR,
				ReadSpeed:  cst.ERROR,
			}
		}
		var readBytes, writeBytes uint64
		for _, v := range counters {
			readBytes += v.ReadBytes
			writeBytes += v.WriteBytes
		}
		LastReadBytes = readBytes
		LastWriteBytes = writeBytes
		time.Sleep(time.Second)
	}

	counters, err := disk.IOCounters()
	if err != nil {
		log.Error(err)
		return DiskRWSpeed{
			WriteSpeed: cst.ERROR,
			ReadSpeed:  cst.ERROR,
		}
	}
	LastGetTime = time.Now().Unix()

	var readBytes, writeBytes uint64
	for _, v := range counters {
		readBytes += v.ReadBytes
		writeBytes += v.WriteBytes
	}

	difWriteBytes := writeBytes - LastWriteBytes
	difReadBytes := readBytes - LastReadBytes
	LastWriteBytes = writeBytes
	LastReadBytes = readBytes

	readSpeed, writeSpeed := "", ""

	difReadBytes = difReadBytes / 1024
	difWriteBytes = difWriteBytes / 1024

	if difReadBytes > 1024 {
		readSpeed = fmt.Sprintf("%.2fMb/s", float64(difReadBytes/1024))
	} else {
		readSpeed = fmt.Sprintf("%dKb/s", difReadBytes)
	}

	if difWriteBytes > 1024 {
		writeSpeed = fmt.Sprintf("%.2fMb/s", float64(difWriteBytes/1024))
	} else {
		writeSpeed = fmt.Sprintf("%dKb/s", difWriteBytes)
	}
	return DiskRWSpeed{
		WriteSpeed: writeSpeed,
		ReadSpeed:  readSpeed,
	}

}
