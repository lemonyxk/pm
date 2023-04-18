/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-17 01:23
**/

package main

import (
	"fmt"
	"os"

	"github.com/kardianos/service"
	"github.com/lemonyxk/console"
	"github.com/lemonyxk/pm/config"
	"github.com/lemonyxk/pm/program"
	"github.com/lemonyxk/pm/system"
)

func init() {

	if !system.IsAdmin() {
		fmt.Println("pm service manager must run as administrator")
		os.Exit(1)
	}

	_ = os.MkdirAll(config.HomeDir, os.ModePerm)
	_ = os.MkdirAll(config.ConfigDir, os.ModePerm)
	_ = os.MkdirAll(config.UnActiveDir, os.ModePerm)
	_ = os.MkdirAll(config.LogDir, os.ModePerm)
	_ = os.MkdirAll(config.VarDir, os.ModePerm)

	var err error
	config.OutFile, err = os.OpenFile(config.OutPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config.ErrFile, err = os.OpenFile(config.ErrPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	console.DefaultLogger.Stdout = config.OutFile
	console.DefaultLogger.Stderr = config.ErrFile

	config.InitConfig()

	var svcConfig = &service.Config{
		Name:        "PM",
		DisplayName: "PM Service",
		Description: "PM Service Manager",
		Arguments:   []string{},
		EnvVars: map[string]string{
			"RUN": "TRUE",
		},

		// Executable:  "",
	}

	prg := &program.Program{}

	config.Server, err = service.New(prg, svcConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
