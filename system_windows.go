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
	"unsafe"

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
	return &syscall.SysProcAttr{HideWindow: true}, nil
}

func handlerCmd(cmd *exec.Cmd) chan struct{} {

	var ch = make(chan struct{}, 1)

	job, err := windows.CreateJobObject(nil, nil)
	if err != nil {
		panic(err)
	}

	go func() {
		<-ch
		windows.CloseHandle(job)
	}()

	info := windows.JOBOBJECT_EXTENDED_LIMIT_INFORMATION{
		BasicLimitInformation: windows.JOBOBJECT_BASIC_LIMIT_INFORMATION{
			LimitFlags: windows.JOB_OBJECT_LIMIT_KILL_ON_JOB_CLOSE,
		},
	}
	if _, err := windows.SetInformationJobObject(
		job,
		windows.JobObjectExtendedLimitInformation,
		uintptr(unsafe.Pointer(&info)),
		uint32(unsafe.Sizeof(info))); err != nil {
		panic(err)
	}

	type process struct {
		Pid    int
		Handle uintptr
	}

	if err := windows.AssignProcessToJobObject(
		job,
		windows.Handle((*process)(unsafe.Pointer(cmd.Process)).Handle)); err != nil {
		panic(err)
	}

	return ch
}
