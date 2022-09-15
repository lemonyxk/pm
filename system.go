//go:build !windows
// +build !windows

/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-14 23:14
**/

package main

import (
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"syscall"

	"github.com/lemonyxk/console"
)

func getPid() int {
	var ps = findProcessByPID(int32(os.Getpid()))
	if len(ps) == 0 {
		console.Exit("can not find process by pid", os.Getpid())
	}

	var p = ps[0]

	for {
		pp, err := p.Parent()
		if err != nil {
			break
		}

		n, err := p.Name()
		if err != nil {
			break
		}

		if strings.ToUpper(n) == "SUDO" {
			break
		}

		p = pp
	}

	return int(p.Pid)
}

func isAdmin() bool {
	return os.Geteuid() == 0
}

func newCmd(command string) *exec.Cmd {
	return exec.Command("/bin/bash", "-c", command)
}

func getSysProcAttr(userName string) (*syscall.SysProcAttr, error) {
	u, err := user.Lookup(userName)
	if err != nil {
		return nil, err
	}

	uid, err := strconv.ParseInt(u.Uid, 10, 32)
	if err != nil {
		return nil, err
	}

	gid, err := strconv.ParseInt(u.Gid, 10, 32)
	if err != nil {
		return nil, err
	}

	return &syscall.SysProcAttr{
		Credential: &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid)},
	}, nil
}

func handlerCmd(cmd *exec.Cmd) chan struct{} {
	var ch = make(chan struct{}, 1)
	go func() {
		<-ch
	}()
	return ch
}
