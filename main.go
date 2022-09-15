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
	"path/filepath"
	"time"

	"github.com/hpcloud/tail"
	"github.com/kardianos/service"
	"github.com/lemonyxk/console"
	hash "github.com/lemonyxk/structure/v3/map"
)

var server service.Service
var homeDir = filepath.Join(homePath(), "pm")
var configDir = filepath.Join(homeDir, "config")
var varDir = filepath.Join(homeDir, "var")
var logDir = filepath.Join(homeDir, "log")
var outPath = filepath.Join(logDir, "out.log")
var errPath = filepath.Join(logDir, "err.log")
var runtimePath = filepath.Join(varDir, "runtime")

var outFile *os.File
var errFile *os.File

var sigMap = hash.NewSync[string, *Proc]()
var config []Config

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() {
	// Do work here
	run()
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	_, err := httpGet("closeChan", nil)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func init() {
	console.SetFlags(0)
	console.Colorful(false)

	_ = os.MkdirAll(homeDir, os.ModePerm)
	_ = os.MkdirAll(configDir, os.ModePerm)
	_ = os.MkdirAll(logDir, os.ModePerm)
	_ = os.MkdirAll(varDir, os.ModePerm)

	var err error
	outFile, err = os.OpenFile(outPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	errFile, err = os.OpenFile(errPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	console.SetWriter(outFile)
	console.SetErrorWriter(errFile)

	initConfig()

	var svcConfig = &service.Config{
		Name:        "PM",
		DisplayName: "PM Service",
		Description: "PM Service Manager",
		Arguments:   []string{"run", "--force"},
		// Executable:  "",
	}

	prg := &program{}

	server, err = service.New(prg, svcConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(Help())
		os.Exit(1)
	}

	if !isAdmin() {
		fmt.Println("pm service manager must run as administrator")
		os.Exit(1)
	}

	switch Args(1) {
	case "install":
		var err = server.Install()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("install success")
		}
	case "uninstall":
		var err = server.Uninstall()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("uninstall success")
		}
	case "start":
		var err = server.Start()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("start success")
		}
	case "stop":
		var err = server.Stop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("stop success")
		}
	case "restart":
		var err = server.Stop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("stop success")
		}
		time.Sleep(time.Millisecond * 100)
		err = server.Start()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("start success")
		}
	case "status":
		var status, err = server.Status()
		if err != nil {
			fmt.Println(err)
		} else {
			switch status {
			case 0:
				fmt.Println("unknown")
			case 1:
				fmt.Println("running")
			case 2:
				fmt.Println("stopped")
			}
		}
	case "run":
		var err = server.Run()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("run success")
		}
	case "log":
		t, err := tail.TailFile(outPath, tail.Config{Follow: true, Poll: true})
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		for line := range t.Lines {
			fmt.Println(line.Text)
		}

	case "err":
		t, err := tail.TailFile(errPath, tail.Config{Follow: true, Poll: true})
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		for line := range t.Lines {
			fmt.Println(line.Text)
		}
	case "info":
		fmt.Println("home dir:", homeDir)
		fmt.Println("config dir:", configDir)
		fmt.Println("out log path:", outPath)
		fmt.Println("err log path:", errPath)
	case "services":
		services()
	default:
		fmt.Println(Help())
	}
}
