//go:build windows
// +build windows

/**
* @program: pm
*
* @description:
*
* @author: lemon
*
* @create: 2022-09-12 18:24
**/

package config

import (
	"os"
	"path/filepath"
)

func HomePath() string {
	return filepath.Join(os.Getenv("PROGRAMDATA"), "opt")
}
