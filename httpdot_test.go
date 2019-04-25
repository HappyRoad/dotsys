package dotsys

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestHttpDot(t *testing.T) {
	go func() {
		http.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
			result, _ := ioutil.ReadAll(r.Body)
			r.Body.Close()
			log.Printf("%s", result)
		})

		http.ListenAndServe(":7777", nil)
	}()

	time.Sleep(time.Second * 3)

	dot, err := NewHttpDot("test")
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		for {
			dot.Add("f1", 1)
			dot.Add("f2", 1)
			time.Sleep(time.Microsecond)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			dot.Add("f1", 1)
			time.Sleep(time.Microsecond * 2)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			dot.Add("f2", 1)
			time.Sleep(time.Microsecond * 3)
		}
	}()

	wg.Wait()
}
