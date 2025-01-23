package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"syscall"
	"time"
)

type Root struct {
	Services  []Service  `json:"services"`
	Sequences []Sequence `json:"sqc"`
	Processes []Process  `json:"ps"`
}

type Service struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Command   string    `json:"command"`
	Path      string    `json:"path"`
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

func new() *Root {
	return &Root{
		Services:  []Service{},
		Sequences: []Sequence{},
		Processes: []Process{},
	}
}

func (r *Root) load() {
	raw, err := os.ReadFile(os.Getenv(ZOITY_PATH_CONFIG) + "/settings.zoity")
	if err != nil {
		fmt.Println(err)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(string(raw))
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = json.Unmarshal(decoded, &r); err != nil {
		fmt.Println(err)
	}
}

func (r *Root) update() {
	bytes, err := json.Marshal(&r)
	if err != nil {
		fmt.Println(err)
		return
	}
	encoded := base64.StdEncoding.EncodeToString(bytes)
	os.WriteFile(os.Getenv(ZOITY_PATH_CONFIG)+"/settings.zoity", []byte(encoded), os.ModePerm)
}

func (r *Root) add(pid int, sid string) {
	r.Processes = append(r.Processes, Process{Pid: pid, Sid: sid, CreatedAt: time.Now()})
}

func (r *Root) kill(sid string) {
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
