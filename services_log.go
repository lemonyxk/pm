/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-15 17:17
**/

package main

import (
	"fmt"

	"github.com/hpcloud/tail"
)

func errService() {
	var name = GetArgs([]string{"err"})
	if name == "" {
		fmt.Println(Help())
		return
	}

	var cfg = getConfigByName(name)
	if cfg.Name == "" {
		fmt.Println("config not found")
		return
	}

	var out, err = getErrPath(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	t, err := tail.TailFile(out, tail.Config{Follow: true, Poll: true})
	ExitIfError(err)

	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}

func logService() {
	var name = GetArgs([]string{"log"})
	if name == "" {
		fmt.Println(Help())
		return
	}

	var cfg = getConfigByName(name)
	if cfg.Name == "" {
		fmt.Println("config not found")
		return
	}

	var out, err = getOutPath(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	t, err := tail.TailFile(out, tail.Config{Follow: true, Poll: true})
	ExitIfError(err)

	for line := range t.Lines {
		fmt.Println(line.Text)
	}
}
