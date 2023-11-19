package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall"
	"time"
)

type Root struct {
	Services  []Service  `json:"services"`
	Sequences []Sequence `json:"sequences"`
	Processes []Process  `json:"processes"`
}

func (r *Root) addProcess(pid int, sid string) {
	r.Processes = append(r.Processes, Process{Pid: pid, Sid: sid, CreatedAt: time.Now()})
}

func (r *Root) killProcess(sid string) {
	for i := range r.Processes {
		if r.Processes[i].Sid == sid {
			syscall.Kill(-r.Processes[i].Pid, syscall.SIGKILL)
			// FIXME: (@isaqueveras) delete item of list
			// r.Processes = slices.Delete(r.Processes, i, i+1)
		}
	}
}

func (r *Root) searchServiceByName(name string) *Service {
	for _, service := range r.Services {
		if service.Name == name {
			return &service
		}
	}
	return nil
}

type Service struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Command   string    `json:"command"`
	Path      string    `json:"path"`
	Port      string    `json:"port"`
	CreatedAt time.Time `json:"created_at"`
}

type Sequence struct {
	Name     string   `json:"name"`
	Services []string `json:"services"`
}

type Process struct {
	Pid       int       `json:"pid"`
	Sid       string    `json:"sid"`
	CreatedAt time.Time `json:"created_at"`
}

func getConfig() (config *Root) {
	raw, err := os.ReadFile(pathRoot + "/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = json.Unmarshal(raw, &config); err != nil {
		fmt.Println(err)
		return
	}

	return config
}

func updateConfig(cfg *Root) error {
	bytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(pathRoot+"/config.json", bytes, os.ModePerm)
}
