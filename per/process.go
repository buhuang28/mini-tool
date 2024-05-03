package per

import (
	"github.com/shirou/gopsutil/v4/process"
	log "github.com/sirupsen/logrus"
)

type Process struct {
	Pid  int32
	Name string
	Path string
}

func GetProcess() []Process {
	processes, err := process.Processes()
	if err != nil {
		log.Error(err)
		return nil
	}

	var processList []Process

	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			log.Error(err)
			continue
		}
		exe, err := p.Exe()
		if err != nil {
			log.Error(err)
			continue
		}
		pid := p.Pid
		processList = append(processList, Process{
			Pid:  pid,
			Name: name,
			Path: exe,
		})
	}
	return processList
}
