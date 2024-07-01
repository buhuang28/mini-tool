package per

import (
	"encoding/json"
	"github.com/buhuang28/mini-tool/caches"
	"github.com/buhuang28/mini-tool/config"
	"github.com/buhuang28/mini-tool/utils"
	"github.com/shirou/gopsutil/v4/process"
	log "github.com/sirupsen/logrus"
)

func GetProcess() []config.Process {
	processes, err := process.Processes()
	if err != nil {
		log.Error(err)
		return nil
	}

	var processList []config.Process

	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			continue
		}
		exe, err := p.Exe()
		if err != nil {
			//log.Error(err)
			continue
		}
		pid := p.Pid
		processList = append(processList, config.Process{
			Pid:  pid,
			Name: name,
			Path: exe,
		})
	}
	return processList
}

func KillProcess() {
	config.CfgLock.Lock()
	defer config.CfgLock.Unlock()

	cfgBytes := utils.ReadFile(caches.ConfigPath)
	if len(cfgBytes) == 0 {
		return
	}
	c := new(config.Config)
	err := json.Unmarshal(cfgBytes, c)
	if err != nil {
		log.Error(err)
		return
	}
	processes, err := process.Processes()
	if err != nil {
		log.Error(err)
		return
	}
	for _, v := range processes {
		processName, err := v.Name()
		if err != nil {
			continue
		}

		for _, v2 := range c.KillName {
			if processName == v2 {
				log.Infof("kill process:%s", v2)
				err = v.Kill()
				if err != nil {
					log.Error()
				}
			}
		}
	}

}
