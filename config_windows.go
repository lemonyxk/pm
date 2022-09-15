//go:build windows
// +build windows

/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-12 18:24
**/

package main

import (
	"os"
)

func homePath() string {
	return os.Getenv("PROGRAMDATA")
}
