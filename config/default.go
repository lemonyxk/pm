/**
* @program: pm
*
* @description:
*
* @author: lemon
*
* @create: 2022-09-16 23:30
**/

package config

import (
	"os"
	"path/filepath"
)

var HomeDir = filepath.Join(HomePath(), "pm")
var CfgDir = filepath.Join(HomeDir, "config")
var VarDir = filepath.Join(HomeDir, "var")
var LogDir = filepath.Join(HomeDir, "log")
var OutPath = filepath.Join(LogDir, "out.log")
var ErrPath = filepath.Join(LogDir, "err.log")

var OutFile *os.File
var ErrFile *os.File
