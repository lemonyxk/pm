/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-12 16:23
**/

package main

import (
	"os"
	"sync/atomic"
	"time"

	"github.com/lemonyxk/console"
)

type Config struct {
	Name    string   `json:"name"`
	User    string   `json:"user"`
	Dir     string   `json:"dir"`
	Command []string `json:"command"`
	Restart bool     `json:"restart"`
	Out     string   `json:"out"`
	Err     string   `json:"err"`
}

func findRuntime(cfg Config, runtime Process) *Proc {
	for i := 0; i < len(runtime); i++ {
		if runtime[i].Name == cfg.Name {
			runtime[i].ch = make(chan struct{}, 1)
			return runtime[i]
		}
	}
	return nil
}

func run() {

	var err = createServer()
	if err != nil {
		console.Exit(err)
	}

	console.Info("pm manager http server start addr:", "127.0.0.1:52525")
	console.Info("pm manager service start pid:", os.Getpid())
	console.Info("home dir:", homeDir)
	console.Info("config dir:", configDir)
	console.Info("out log path:", outPath)
	console.Info("err log path:", errPath)

	for i := 0; i < len(config); i++ {
		go start(config[i])
	}

	<-closeChan

	console.Info("pm manager service stopped")
}

func start(cfg Config) {
	if len(cfg.Command) == 0 {
		return
	}

	var cInfo, err = makeCmdInfo(cfg)
	if err != nil {
		console.Error(err)
		return
	}

	var proc = &Proc{
		Name:    cfg.Name,
		Cmd:     cfg.Command,
		OutPath: cInfo.outPath,
		ErrPath: cInfo.errPath,

		Children: nil,

		ch: make(chan struct{}, 1),
	}

	var fin int32 = 0
	sigMap.Set(cfg.Name, proc)

	for j := 0; j < len(cfg.Command); j++ {
		var cmdS = cfg.Command[j]
		var child = &Child{Pid: 0, Restart: cfg.Restart, Status: "stop"}
		proc.Children = append(proc.Children, child)

		go func() {
			for {

				var cmd = newCmd(cmdS)
				cmd.Dir = cInfo.dir
				cmd.SysProcAttr = cInfo.procAttr
				cmd.Stdout = cInfo.stdout
				cmd.Stderr = cInfo.stderr
				cmd.Stdin = os.Stdin

				err = cmd.Start()
				if err != nil {
					console.Error(err)
					time.Sleep(time.Second)
					break
				}

				child.Pid = cmd.Process.Pid
				child.Status = "running"

				console.Info("start process", cfg.Name, "pid is", cmd.Process.Pid)

				var ch = handlerCmd(cmd)

				err = cmd.Wait()
				if err != nil {
					console.Error(err)
				}

				ch <- struct{}{}

				child.Status = "stop"
				child.Pid = 0

				if !child.Restart {
					console.Info("stop process", cfg.Name, "pid is", cmd.Process.Pid)
					if atomic.AddInt32(&fin, 1) == int32(len(cfg.Command)) {
						proc.ch <- struct{}{}
					}
					break
				}

				console.Info("stop process", cfg.Name, "pid is", cmd.Process.Pid, "trying to restart...")
				time.Sleep(time.Second)
			}
		}()
	}

	<-proc.ch

	sigMap.Delete(cfg.Name)

	if cInfo.of != nil {
		_ = cInfo.of.Close()
	}

	if cInfo.ef != nil {
		_ = cInfo.ef.Close()
	}

	console.Info("service", cfg.Name, "stopped")
}
