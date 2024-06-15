package config

import (
	"encoding/json"
	"github.com/buhuang28/mini-tool/caches"
	"github.com/buhuang28/mini-tool/etype"
	"github.com/buhuang28/mini-tool/utils"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	KillName     []string       `json:"kill_name,omitempty"`
	LimitNetPath []LimitNetInfo `json:"limit_net_path,omitempty"`
}

type LimitNetInfo struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

var (
	CfgLock = new(sync.Mutex)
)

func AddKillName(name string) error {
	CfgLock.Lock()
	defer CfgLock.Unlock()
	cfgBytes := utils.ReadFile(caches.ConfigPath)
	c := new(Config)
	if len(cfgBytes) != 0 {
		err := json.Unmarshal(cfgBytes, c)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	for _, v := range c.KillName {
		if v == name {
			return etype.NAME_EXIST
		}
	}
	c.KillName = append(c.KillName, name)
	marshal, err := json.Marshal(c)
	if err != nil {
		log.Error(err)
		return err
	}
	utils.WriteFile(caches.ConfigPath, marshal)
	return nil
}
