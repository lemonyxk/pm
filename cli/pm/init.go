/**
* @program: pm
*
* @description:
*
* @author: lemon
*
* @create: 2022-09-17 01:23
**/

package main

import (
	"fmt"
	"path/filepath"

	"github.com/hpcloud/tail"
	"github.com/lemonyxk/pm/config"
	"github.com/lemonyxk/pm/process"
	"github.com/lemonyxk/pm/tools"
	"github.com/lemonyxk/utils/v3"
)

func init() {

}

func GetService(name string) *process.Proc {
	var ss = GetServices()
	for i := 0; i < len(ss); i++ {
		if ss[i].Name == name {
			return ss[i]
		}
	}
	return nil
}

func GetServices() process.Process {
	var list process.Process
	body, err := tools.HttpGet("list", nil)
	tools.ExitIfError(err)
	_ = utils.Json.Decode(body, &list)
	return list
}

func Query(action string) string {
	var name = tools.GetArgs([]string{action})
	var bts, err = tools.HttpGet(action, tools.M{"name": name})
	if err != nil {
		return err.Error()
	}
	return string(bts)
}

func errService() {
	var name = tools.GetArgs([]string{"err"})
	if name == "" {
		fmt.Println(tools.PMHelp())
		return
	}

	var cfgPath = filepath.Join(config.CfgDir, name+".json")
	var res = utils.File.ReadFromPath(cfgPath)
	if res.LastError() != nil {
		fmt.Println(res.LastError())
		return
	}

	var cfg config.Config
	var err = utils.Json.Decode(res.Bytes(), &cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(cfg.Command) == 0 {
		fmt.Println("config not found")
		return
	}

	cfg.Name = name

	out, err := config.GetErrPath(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	t, err := tail.TailFile(out, tail.Config{Follow: true, Poll: true})
	tools.ExitIfError(err)

	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}

func logService() {
	var name = tools.GetArgs([]string{"log"})
	if name == "" {
		fmt.Println(tools.PMHelp())
		return
	}

	var cfgPath = filepath.Join(config.CfgDir, name+".json")
	var res = utils.File.ReadFromPath(cfgPath)
	if res.LastError() != nil {
		fmt.Println(res.LastError())
		return
	}

	var cfg config.Config
	var err = utils.Json.Decode(res.Bytes(), &cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(cfg.Command) == 0 {
		fmt.Println("config not found")
		return
	}

	cfg.Name = name

	out, err := config.GetOutPath(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	t, err := tail.TailFile(out, tail.Config{Follow: true, Poll: true})
	tools.ExitIfError(err)

	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
