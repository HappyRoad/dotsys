package dotsys

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	DotServerURL = "http://dotsys.xtgreat.com:7777/heartbeat"
)

func (r Request) ToBuffer() *bytes.Buffer {
	data, _ := json.Marshal(r)
	return bytes.NewBuffer(data)
}

type HttpHeartBeat struct {
	client *http.Client
	dot    *Dot
	name   string
}

func NewHttpDot(name string) (dot *Dot, err error) {
	hb := &HttpHeartBeat{
		client: &http.Client{
			Timeout: time.Millisecond * 30,
		},
		name:   name,
	}
	hb.dot, _ = newDot(hb)

	return hb.dot, nil
}

func (hb HttpHeartBeat) HeartBeat() error {
	request := &Request{
		Host:      HostName,
		Ip:        Ip,
		Project:      hb.name,
		DotValues: hb.dot.getDotValues(),
	}
	res, err := hb.client.Post(DotServerURL, "application/json", request.ToBuffer())
	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()
	if err != nil {
		return errors.New(fmt.Sprintf("http heartbeat error: post error %s", err.Error()))
	}
	if res == nil {
		return errors.New("http heartbeat error: server no response")
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("http heartbeat error: bad status %d", res.StatusCode))
	}
	return nil
}
