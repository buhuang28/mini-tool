package per

import (
	"fmt"
	"github.com/buhuang28/mini-tool/cst"
	"github.com/shirou/gopsutil/v4/disk"
	log "github.com/sirupsen/logrus"
	"time"
)

type DiskRWSpeed struct {
	WriteSpeed string
	ReadSpeed  string
}

func GetDiskWRSpeed() DiskRWSpeed {
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
	time.Sleep(time.Second)

	counters, err = disk.IOCounters()
	if err != nil {
		log.Error(err)
		return DiskRWSpeed{
			WriteSpeed: cst.ERROR,
			ReadSpeed:  cst.ERROR,
		}
	}

	var readBytes2, writeBytes2 uint64
	for _, v := range counters {
		readBytes2 += v.ReadBytes
		writeBytes2 += v.WriteBytes
		//readBytes = v.ReadBytes-readBytes
		//writeBytes = v.WriteBytes - writeBytes
	}

	writeBytes = writeBytes2 - writeBytes
	readBytes = readBytes2 - readBytes

	readSpeed, writeSpeed := "", ""

	readBytes = readBytes / 1024
	writeBytes = writeBytes / 1024

	if readBytes > 1024 {
		readSpeed = fmt.Sprintf("%.2fMb/s", float64(readBytes/1024))
	} else {
		readSpeed = fmt.Sprintf("%dKb/s", readBytes)
	}

	if writeBytes > 1024 {
		writeSpeed = fmt.Sprintf("%.2fMb/s", float64(writeBytes/1024))
	} else {
		writeSpeed = fmt.Sprintf("%dKb/s", writeBytes)
	}
	return DiskRWSpeed{
		WriteSpeed: writeSpeed,
		ReadSpeed:  readSpeed,
	}

}
