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

package config

import (
	"os"
)

func HomePath() string {
	return os.Getenv("PROGRAMDATA")
}
