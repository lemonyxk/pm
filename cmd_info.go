/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-15 16:27
**/

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"
)

type cmdInfo struct {
	of             *os.File
	ef             *os.File
	stdout, stderr io.Writer
	procAttr       *syscall.SysProcAttr
	dir            string

	outPath string
	errPath string
}

func makeCmdInfo(cfg Config) (*cmdInfo, error) {

	var cmdInfo = &cmdInfo{}

	if cfg.User != "" {
		var c, err = getSysProcAttr(cfg.User)
		if err != nil {
			return nil, err
		}
		cmdInfo.procAttr = c
	}

	if cfg.Dir != "" {
		var dir, err = filepath.Abs(cfg.Dir)
		if err != nil {
			return nil, err
		}

		cmdInfo.dir = dir
	}

	{
		out, err := getOutPath(cfg)
		if err != nil {
			return nil, err
		}

		of, err := os.OpenFile(out, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err != nil {
			return nil, err
		}

		cmdInfo.stdout = of
		cmdInfo.of = of
		cmdInfo.outPath = out
	}

	{
		out, err := getErrPath(cfg)
		if err != nil {
			return nil, err
		}

		ef, err := os.OpenFile(out, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err != nil {
			return nil, err
		}

		cmdInfo.stderr = ef
		cmdInfo.ef = ef
		cmdInfo.errPath = out
	}

	return cmdInfo, nil
}

func getErrPath(cfg Config) (string, error) {
	var out = cfg.Err

	if cfg.Err == "" {
		out = filepath.Join(logDir, cfg.Name+".err.log")
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

func getOutPath(cfg Config) (string, error) {
	var out = cfg.Out

	if cfg.Out == "" {
		out = filepath.Join(logDir, cfg.Name+".out.log")
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
