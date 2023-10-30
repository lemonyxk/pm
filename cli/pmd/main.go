/**
* @program: pm
*
* @description:
*
* @author: lemon
*
* @create: 2022-09-12 16:13
**/

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kardianos/service"
	"github.com/lemonyxk/console"
	"github.com/lemonyxk/pm/app"
	"github.com/lemonyxk/pm/config"
	"github.com/lemonyxk/pm/tools"
	"github.com/nxadm/tail"
)

func main() {
	console.Info("pm service manager", os.Environ())
	if os.Getenv("RUN") == "TRUE" {
		var err = app.Server.Run()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("run success")
		}
	}

	switch tools.Args(1) {
	case "install":
		// install main process
		_ = app.Server.Install()
	case "uninstall":
		var err = app.Server.Stop()
		if err != nil {
			fmt.Println("already stopped")
		} else {
			fmt.Println("stop success")
		}
		_ = app.Server.Uninstall()
	case "start":
		var err = app.Server.Start()
		if err != nil {
			fmt.Println("already started")
		} else {
			fmt.Println("start success")
		}
	case "stop":
		// stop main process
		var err = app.Server.Stop()
		if err != nil {
			fmt.Println("already stopped")
		} else {
			fmt.Println("stop success")
		}
	case "restart":
		var err = app.Server.Stop()
		if err != nil {
			fmt.Println("already stopped")
		} else {
			fmt.Println("stop success")
		}
		for {
			time.Sleep(time.Second * 1)
			var status, err = app.Server.Status()
			if err != nil {
				fmt.Println(err)
				break
			}
			if status == service.StatusStopped {
				break
			}
		}
		err = app.Server.Start()
		if err != nil {
			fmt.Println("already started")
		} else {
			fmt.Println("start success")
		}
	case "status":
		var status, err = app.Server.Status()
		if err != nil {
			fmt.Println(err)
		} else {
			switch status {
			case service.StatusUnknown:
				fmt.Println("unknown")
			case service.StatusRunning:
				fmt.Println("running")
			case service.StatusStopped:
				fmt.Println("stopped")
			}
		}
	case "log":
		t, err := tail.TailFile(config.OutPath, tail.Config{Follow: true, Poll: true})
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		for line := range t.Lines {
			fmt.Println(line.Text)
		}

	case "err":
		t, err := tail.TailFile(config.ErrPath, tail.Config{Follow: true, Poll: true})
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		for line := range t.Lines {
			fmt.Println(line.Text)
		}
	case "info":
		fmt.Println("home:", config.HomeDir)
		fmt.Println("config:", config.CfgDir)
		fmt.Println("var:", config.VarDir)
		fmt.Println("log:", config.LogDir)
	case "":
		fmt.Println(tools.PMDHelp())
	default:
		fmt.Println(tools.PMDHelp())
	}
}
