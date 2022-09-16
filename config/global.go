/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-16 23:30
**/

package config

import (
	"os"
	"path/filepath"

	"github.com/kardianos/service"
	"github.com/lemonyxk/pm/process"
	hash "github.com/lemonyxk/structure/v3/map"
)

var Server service.Service
var HomeDir = filepath.Join(HomePath(), "pm")
var ConfigDir = filepath.Join(HomeDir, "config")
var UnActiveDir = filepath.Join(HomeDir, "unActive")
var VarDir = filepath.Join(HomeDir, "var")
var LogDir = filepath.Join(HomeDir, "log")
var OutPath = filepath.Join(LogDir, "out.log")
var ErrPath = filepath.Join(LogDir, "err.log")
var RuntimePath = filepath.Join(VarDir, "runtime")

var OutFile *os.File
var ErrFile *os.File

var SigMap = hash.NewSync[string, *process.Proc]()
