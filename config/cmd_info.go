/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-15 16:27
**/

package config

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/lemonyxk/pm/system"
)

type CmdInfo struct {
	OutFile     *os.File
	ErrFile     *os.File
	SysProcAttr *syscall.SysProcAttr
	Dir         string

	OutPath string
	ErrPath string
}

func NewCmdInfo(cfg Config) (*CmdInfo, error) {

	var cmdInfo = &CmdInfo{}

	if cfg.User != "" {
		var c, err = system.GetSysProcAttr(cfg.User)
		if err != nil {
			return nil, err
		}
		cmdInfo.SysProcAttr = c
	}

	if cfg.Dir != "" {
		var dir, err = filepath.Abs(cfg.Dir)
		if err != nil {
			return nil, err
		}

		cmdInfo.Dir = dir
	}

	{
		out, err := GetOutPath(cfg)
		if err != nil {
			return nil, err
		}

		of, err := os.OpenFile(out, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err != nil {
			return nil, err
		}

		cmdInfo.OutFile = of
		cmdInfo.OutPath = out
	}

	{
		out, err := GetErrPath(cfg)
		if err != nil {
			return nil, err
		}

		ef, err := os.OpenFile(out, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err != nil {
			return nil, err
		}

		cmdInfo.ErrFile = ef
		cmdInfo.ErrPath = out
	}

	return cmdInfo, nil
}

func GetErrPath(cfg Config) (string, error) {
	var out = cfg.Err

	if cfg.Err == "" {
		out = filepath.Join(LogDir, cfg.Name+".err.log")
	}

	if !filepath.IsAbs(out) {
		if cfg.Dir == "" {
			return "", fmt.Errorf("%s is not an absolute path", out)
		}
		out = filepath.Join(cfg.Dir, out)
	}

	out, err := filepath.Abs(out)
	if err != nil {
		return "", err
	}

	return out, nil
}

func GetOutPath(cfg Config) (string, error) {
	var out = cfg.Out

	if cfg.Out == "" {
		out = filepath.Join(LogDir, cfg.Name+".out.log")
	}

	if !filepath.IsAbs(out) {
		if cfg.Dir == "" {
			return "", fmt.Errorf("%s is not an absolute path", out)
		}
		out = filepath.Join(cfg.Dir, out)
	}

	out, err := filepath.Abs(out)
	if err != nil {
		return "", err
	}

	return out, nil
}
