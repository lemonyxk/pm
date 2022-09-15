//go:build windows
// +build windows

/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-14 23:15
**/

package main

import (
	"os"
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"
)

func getPid() int {
	return os.Getpid()
}

func isAdmin() bool {
	// Directly copied from the official windows documentation. The Go API for this is a
	// direct wrap around the official C++ API.
	// See https://docs.microsoft.com/en-us/windows/desktop/api/securitybaseapi/nf-securitybaseapi-checktokenmembership
	var sid *windows.SID
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid)
	if err != nil {
		return false
	}
	adm, err := windows.GetCurrentProcessToken().IsMember(sid)
	return err == nil && adm
}

func newCmd(command string) *exec.Cmd {
	return exec.Command("cmd", "/c", command)
}

func getSysProcAttr(userName string) (*syscall.SysProcAttr, error) {
	return nil, nil
}
