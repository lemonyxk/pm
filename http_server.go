/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-15 03:03
**/

package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/lemonyxk/console"
	"github.com/lemonyxk/utils/v3"
)

var closeChan = make(chan struct{}, 1)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		h.list(w, r)
	case "/stop":
		h.stop(w, r)
	case "/stopAll":
		h.stopAll(w, r)
	case "/start":
		h.start(w, r)
	case "/restart":
		h.restart(w, r)
	case "/closeChan":
		closeChan <- struct{}{}
		h.endStr(w, nil)
	default:
		http.NotFound(w, r)
	}
}

func (h *handler) list(w http.ResponseWriter, r *http.Request) {
	var list Process

	sigMap.Range(func(k string, v *Proc) bool {
		list = append(list, v)
		return true
	})

	h.end(w, list)
}

func (h *handler) stopAll(w http.ResponseWriter, r *http.Request) {

	var str = ""

	sigMap.Range(func(k string, v *Proc) bool {
		for i := 0; i < len(v.Children); i++ {
			v.Children[i].Restart = false
			var p = findProcessByPID(int32(v.Children[i].Pid))
			if len(p) == 0 {
				continue
			}
			_ = p[0].Kill()
			str += fmt.Sprintf("kill process %d", v.Children[i].Pid) + "\n"
		}

		str += "stop " + k + " success\n"

		return true
	})

	h.endStr(w, str+"stop all success")
}

func (h *handler) stop(w http.ResponseWriter, r *http.Request) {
	var q = r.URL.Query()
	var name = q.Get("name")
	if name == "" {
		h.endStr(w, "name is empty")
		return
	}

	var m = sigMap.Get(name)
	if m == nil {
		h.endStr(w, fmt.Sprintf("service %s is not running", name))
		return
	}

	var str = ""

	for i := 0; i < len(m.Children); i++ {
		m.Children[i].Restart = false
		var p = findProcessByPID(int32(m.Children[i].Pid))
		if len(p) == 0 {
			continue
		}
		_ = p[0].Kill()
		str += fmt.Sprintf("kill process %d", m.Children[i].Pid) + "\n"
	}

	h.endStr(w, str+"stop success")
}

func (h *handler) start(w http.ResponseWriter, r *http.Request) {
	var q = r.URL.Query()
	var name = q.Get("name")
	if name == "" {
		h.endStr(w, "name is empty")
		return
	}

	var m = sigMap.Get(name)
	if m != nil {
		h.endStr(w, fmt.Sprintf("service %s is running", name))
		return
	}

	var cfg = getConfigByName(name)
	if cfg.Name == "" {
		h.endStr(w, fmt.Sprintf("service %s is not found", name))
		return
	}

	go start(cfg)

	h.endStr(w, "start success")
}

func (h *handler) restart(w http.ResponseWriter, r *http.Request) {
	var q = r.URL.Query()
	var name = q.Get("name")
	if name == "" {
		h.endStr(w, "name is empty")
		return
	}

	var str = ""

	var m = sigMap.Get(name)
	if m == nil {
		str += fmt.Sprintf("service %s is not running", name) + "\n"
	} else {
		for i := 0; i < len(m.Children); i++ {
			m.Children[i].Restart = false
			var p = findProcessByPID(int32(m.Children[i].Pid))
			if len(p) == 0 {
				continue
			}
			_ = p[0].Kill()
			str += fmt.Sprintf("kill process %d", m.Children[i].Pid) + "\n"
		}
	}

	for {
		time.Sleep(time.Millisecond * 100)
		var m = sigMap.Get(name)
		if m == nil {
			break
		}
	}

	var cfg = getConfigByName(name)
	if cfg.Name == "" {
		str += fmt.Sprintf("service %s is not found", name) + "\n"
		return
	}

	go start(cfg)

	h.endStr(w, str+"start success")
}

func (h *handler) end(w http.ResponseWriter, v any) {
	w.WriteHeader(200)
	_, _ = w.Write(utils.Json.Encode(v))
}

func (h *handler) endStr(w http.ResponseWriter, v any) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte(fmt.Sprintf("%v", v)))
}

func createServer() error {
	var err error
	var netListen net.Listener
	var server = http.Server{Addr: ":52525", Handler: &handler{}}

	netListen, err = net.Listen("tcp", server.Addr)

	if err != nil {
		return err
	}

	go func() {
		err = server.Serve(netListen)
		if err != nil {
			console.Info(err)
		}
	}()

	return nil
}