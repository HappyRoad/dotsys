package dotsys

import (
	"log"
	"os"
	"sync"
	"time"
)

var (
	HostName = ""
	Ip       = ""
)

func init() {
	HostName, _ = os.Hostname()
	Ip, _ = getInnerIp()
}

type Request struct {
	Host      string           `json:"host"`
	Ip        string           `json:"ip"`
	Project      string           `json:"project"`
	DotValues map[string]int64 `json:"dot_values"`
}

type IHeartBeat interface {
	HeartBeat() error
}


type Dot struct {
	mutex     sync.Mutex
	dotValues map[string]int64
	hearBeat  IHeartBeat
}

func newDot(hearBeat IHeartBeat) (dot *Dot, err error) {
	dot = &Dot{
		hearBeat:  hearBeat,
		dotValues: make(map[string]int64),
	}

	go dot.start()

	return
}

func (dot *Dot) Add(key string, value int64) {
	dot.mutex.Lock()
	defer dot.mutex.Unlock()

	if _, ok := dot.dotValues[key]; ok {
		dot.dotValues[key] += value
	} else {
		dot.dotValues[key] = value
	}
}

func (dot *Dot) start() {
	var e error = nil

	t := time.NewTicker(time.Second * 3)
	for {
		select {
		case <-t.C:
			if e = dot.hearBeat.HeartBeat(); e != nil {
				log.Printf("dot heartbeat error: %s", e.Error())
			}
		}
	}
}

func (dot *Dot) getDotValues() map[string]int64 {
	dot.mutex.Lock()
	defer dot.mutex.Unlock()

	data := make(map[string]int64)
	for key, value := range dot.dotValues {
		data[key] = value
		dot.dotValues[key] -= data[key]
	}

	return data
}
