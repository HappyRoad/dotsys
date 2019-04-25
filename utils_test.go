package dotsys

import (
	"testing"
)

func TestGetInnerIp(t *testing.T) {
	if ip, err := getInnerIp(); err != nil {
		t.Errorf("error: %s", err.Error())
	} else {
		t.Logf("inner ip: %s", ip)
	}
}
