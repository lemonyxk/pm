/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-12 16:23
**/

package program

import (
	"os"
	"sync/atomic"
	"time"

	"github.com/lemonyxk/console"
	"github.com/lemonyxk/pm/config"
	"github.com/lemonyxk/pm/process"
	"github.com/lemonyxk/pm/system"
)

func run() {

	var err = CreateServer()
	if err != nil {
		console.Exit(err)
	}

	console.Info("pm manager http server start addr:", "127.0.0.1:52525")
	console.Info("pm manager service start pid:", os.Getpid())
	console.Info("home dir:", config.HomeDir)
	console.Info("config dir:", config.ConfigDir)
	console.Info("out log path:", config.OutPath)
	console.Info("err log path:", config.ErrPath)

	for i := 0; i < len(config.Configs); i++ {
		go Exec(config.Configs[i])
	}

	<-system.Exit

	console.Info("pm manager service stopped")
}

func Exec(cfg config.Config) {
	if len(cfg.Command) == 0 {
		return
	}

	var cInfo, err = config.NewCmdInfo(cfg)
	if err != nil {
		console.Error(err)
		return
	}

	var proc = &process.Proc{
		Name:    cfg.Name,
		Cmd:     cfg.Command,
		OutPath: cInfo.OutPath,
		ErrPath: cInfo.ErrPath,

		Children: nil,

		Ch: make(chan struct{}, 1),
	}

	var fin int32 = 0

	config.SigMap.Set(cfg.Name, proc)

	for j := 0; j < len(cfg.Command); j++ {
		var cmdS = cfg.Command[j]
		var child = &process.Child{Pid: 0, Restart: cfg.Restart, Status: "stop"}
		proc.Children = append(proc.Children, child)

		go func() {
			for {

				var cmd = system.NewCmd(cmdS)
				cmd.Dir = cInfo.Dir
				cmd.SysProcAttr = cInfo.SysProcAttr
				cmd.Stdout = cInfo.OutFile
				cmd.Stderr = cInfo.ErrFile
				cmd.Stdin = os.Stdin

				err = cmd.Start()
				if err != nil {
					console.Error(err)
					time.Sleep(time.Second * 3)
					continue
				}

				child.Pid = cmd.Process.Pid
				child.Status = "running"

				console.Info("start process", cfg.Name, "pid is", cmd.Process.Pid)

				var ch = system.HandlerCmd(cmd)

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
						proc.Ch <- struct{}{}
					}
					break
				}

				console.Info("stop process", cfg.Name, "pid is", cmd.Process.Pid, "trying to restart...")

				time.Sleep(time.Second * 3)
			}
		}()
	}

	<-proc.Ch

	config.SigMap.Delete(cfg.Name)

	if cInfo.OutFile != nil {
		_ = cInfo.OutFile.Close()
	}

	if cInfo.ErrFile != nil {
		_ = cInfo.ErrFile.Close()
	}

	console.Info("service", cfg.Name, "stopped")
}
