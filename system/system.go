//go:build !windows
// +build !windows

/**
* @program: pm
*
* @description:
*
* @author: lemon
*
* @create: 2022-09-14 23:14
**/

package system

import (
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

var Exit = make(chan struct{}, 1)

func IsAdmin() bool {
	return os.Geteuid() == 0
}

func NewCmd(command string) *exec.Cmd {
	return exec.Command("/bin/bash", "-c", command)
}

func GetSysProcAttr(userName string) (*syscall.SysProcAttr, error) {
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

func HandlerCmd(cmd *exec.Cmd) chan struct{} {
	var ch = make(chan struct{}, 1)
	go func() {
		<-ch
	}()
	return ch
}
