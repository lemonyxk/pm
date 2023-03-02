/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-15 03:03
**/

package program

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/lemonyxk/console"
	"github.com/lemonyxk/pm/config"
	"github.com/lemonyxk/pm/def"
	"github.com/lemonyxk/pm/process"
	"github.com/lemonyxk/pm/system"
	"github.com/lemonyxk/pm/tools"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	console.Info("get request:", r.URL.String())

	switch tools.FixURL(r.URL.Path) {
	// client command
	case def.LIST:
		h.list(w, r)
	case def.STOP:
		h.stop(w, r)
	case def.STOPALL:
		h.stopAll(w, r)
	case def.START:
		h.start(w, r)
	case def.RESTART:
		h.restart(w, r)
	case def.REMOVE:
		h.remove(w, r)
	case def.ACTIVE:
		h.active(w, r)
	case def.UNACTIVE:
		h.unActive(w, r)
	case def.LOAD:
		h.load(w, r)

	// 	server command
	case def.EXIT:
		system.Exit <- struct{}{}
		h.endStr(w, nil)

	default:
		http.NotFound(w, r)
	}
}

func (h *handler) list(w http.ResponseWriter, r *http.Request) {
	var list process.Process
	config.SigMap.Range(func(k string, v *process.Proc) bool {
		list = append(list, v)
		return true
	})

	h.end(w, list)
}

func (h *handler) stopAll(w http.ResponseWriter, r *http.Request) {

	var str = ""

	config.SigMap.Range(func(k string, v *process.Proc) bool {
		for i := 0; i < len(v.Children); i++ {
			v.Children[i].Restart = false
			var p = tools.FindProcess(int32(v.Children[i].Pid))
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

	var m = config.SigMap.Get(name)
	if m == nil {
		h.endStr(w, fmt.Sprintf("service %s is not running", name))
		return
	}

	var str = ""

	for i := 0; i < len(m.Children); i++ {
		m.Children[i].Restart = false
		var p = tools.FindProcess(int32(m.Children[i].Pid))
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

	var m = config.SigMap.Get(name)
	if m != nil {
		h.endStr(w, fmt.Sprintf("service %s is running", name))
		return
	}

	config.InitConfig()

	var cfg = tools.GetConfig(name)
	if cfg.Name == "" {
		h.endStr(w, fmt.Sprintf("service %s is not found", name))
		return
	}

	go Exec(cfg)

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

	var m = config.SigMap.Get(name)
	if m == nil {
		str += fmt.Sprintf("service %s is not running", name) + "\n"
	} else {
		for i := 0; i < len(m.Children); i++ {
			m.Children[i].Restart = false
			var p = tools.FindProcess(int32(m.Children[i].Pid))
			if len(p) == 0 {
				continue
			}
			_ = p[0].Kill()
			str += fmt.Sprintf("kill process %d", m.Children[i].Pid) + "\n"
		}
	}

	for {
		time.Sleep(time.Second * 1)
		var m = config.SigMap.Get(name)
		if m == nil {
			break
		}
	}

	config.InitConfig()

	var cfg = tools.GetConfig(name)
	if cfg.Name == "" {
		str += fmt.Sprintf("service %s is not found", name) + "\n"
		h.endStr(w, str+"restart fail")
		return
	}

	go Exec(cfg)

	h.endStr(w, str+"start success")
}

func (h *handler) remove(w http.ResponseWriter, r *http.Request) {
	var q = r.URL.Query()
	var name = q.Get("name")
	if name == "" {
		h.endStr(w, "name is empty")
		return
	}

	var str = ""

	var m = config.SigMap.Get(name)
	if m != nil {
		for i := 0; i < len(m.Children); i++ {
			m.Children[i].Restart = false
			var p = tools.FindProcess(int32(m.Children[i].Pid))
			if len(p) == 0 {
				continue
			}
			_ = p[0].Kill()
			str += fmt.Sprintf("kill process %d", m.Children[i].Pid) + "\n"
		}
		str += "stop success\n"
	}

	_ = os.Remove(filepath.Join(config.ConfigDir, name+".json"))
	_ = os.Remove(filepath.Join(config.LogDir, name+".out.log"))
	_ = os.Remove(filepath.Join(config.LogDir, name+".err.log"))

	config.InitConfig()

	h.endStr(w, str+"config remove success")
}

func (h *handler) active(w http.ResponseWriter, r *http.Request) {
	var q = r.URL.Query()
	var name = q.Get("name")
	if name == "" {
		h.endStr(w, "name is empty")
		return
	}

	if _, err := os.Stat(filepath.Join(config.ConfigDir, name+".json")); err == nil {
		h.endStr(w, name+" already active")
		return
	}

	if _, err := os.Stat(filepath.Join(config.UnActiveDir, name+".json")); err != nil {
		h.endStr(w, name+" not found")
		return
	}

	_ = os.Rename(filepath.Join(config.UnActiveDir, name+".json"), filepath.Join(config.ConfigDir, name+".json"))

	config.InitConfig()

	h.endStr(w, name+" active success")
}

func (h *handler) unActive(w http.ResponseWriter, r *http.Request) {
	var q = r.URL.Query()
	var name = q.Get("name")
	if name == "" {
		h.endStr(w, "name is empty")
		return
	}

	var str = ""

	var m = config.SigMap.Get(name)
	if m != nil {
		for i := 0; i < len(m.Children); i++ {
			m.Children[i].Restart = false
			var p = tools.FindProcess(int32(m.Children[i].Pid))
			if len(p) == 0 {
				continue
			}
			_ = p[0].Kill()
			str += fmt.Sprintf("kill process %d", m.Children[i].Pid) + "\n"
		}
		str += "stop success\n"
	}

	if _, err := os.Stat(filepath.Join(config.ConfigDir, name+".json")); err != nil {
		h.endStr(w, str+name+" not found")
		return
	}

	if _, err := os.Stat(filepath.Join(config.UnActiveDir, name+".json")); err == nil {
		h.endStr(w, str+name+" already unActive")
		return
	}

	_ = os.Rename(filepath.Join(config.ConfigDir, name+".json"), filepath.Join(config.UnActiveDir, name+".json"))

	config.InitConfig()

	h.endStr(w, str+name+" unActive success")
}

func (h *handler) load(w http.ResponseWriter, r *http.Request) {
	config.InitConfig()
	h.endStr(w, "load config success")
}
