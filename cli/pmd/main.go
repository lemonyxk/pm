/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-12 16:13
**/

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hpcloud/tail"
	"github.com/kardianos/service"
	"github.com/lemonyxk/pm/config"
	"github.com/lemonyxk/pm/tools"
)

func main() {

	if os.Getenv("RUN") == "TRUE" {
		var err = config.Server.Run()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("run success")
		}
	}

	switch tools.Args(1) {
	// case "install":
	// 	var err = config.Server.Install()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println("install success")
	// 	}
	// case "uninstall":
	// 	var err = config.Server.Uninstall()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println("uninstall success")
	// 	}
	case "start":
		// start main process
		// will exec run function
		_ = config.Server.Install()
		var err = config.Server.Start()
		if err != nil {
			fmt.Println("already started")
		} else {
			fmt.Println("start success")
		}
	case "stop":
		// stop main process
		var err = config.Server.Stop()
		if err != nil {
			fmt.Println("already stopped")
		} else {
			fmt.Println("stop success")
		}
		_ = config.Server.Uninstall()
	case "restart":
		var err = config.Server.Stop()
		if err != nil {
			fmt.Println("already stopped")
		} else {
			fmt.Println("stop success")
		}
		_ = config.Server.Uninstall()
		_ = config.Server.Install()
		for {
			time.Sleep(time.Second * 1)
			var status, err = config.Server.Status()
			if err != nil {
				fmt.Println(err)
				break
			}
			if status == service.StatusStopped {
				break
			}
		}
		err = config.Server.Start()
		if err != nil {
			fmt.Println("already started")
		} else {
			fmt.Println("start success")
		}
	// case "run":
	// 	var err = config.Server.Run()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println("run success")
	// 	}
	case "status":
		var status, err = config.Server.Status()
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
		fmt.Println("home dir:", config.HomeDir)
		fmt.Println("config dir:", config.ConfigDir)
		fmt.Println("unActive dir:", config.UnActiveDir)
		fmt.Println("var dir:", config.VarDir)
		fmt.Println("log dir:", config.LogDir)
	case "":
		fmt.Println(tools.PMDHelp())
	default:
		fmt.Println(tools.PMDHelp())
	}
}
