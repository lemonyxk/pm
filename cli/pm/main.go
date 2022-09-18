/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-12 16:13
**/

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/lemonyxk/pm/def"
	"github.com/lemonyxk/pm/tools"
)

func main() {
	switch tools.Args(1) {
	case def.LIST:
		fmt.Println(GetServices())
	case def.LOG:
		logService()
	case def.ERR:
		errService()
	case def.STOP:
		Query(def.STOP)
	case def.STOPALL:
		Query(def.STOPALL)
	case def.START:
		Query(def.START)
	case def.RESTART:
		Query(def.RESTART)
	case def.REMOVE:
		Query(def.REMOVE)
	case def.ACTIVE:
		Query(def.ACTIVE)
	case def.UNACTIVE:
		Query(def.UNACTIVE)
	case def.LOAD:
		Query(def.LOAD)
	case "server":
		runCmd(tools.Args(2))
	case "":
		fmt.Println(GetServices())
	case "-h", "--help":
		fmt.Println(tools.PMHelp())
	default:
		fmt.Println(tools.PMHelp())
	}
}

func runCmd(str string) {
	var cmd = exec.Command("pmd", str)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	var err = cmd.Start()
	tools.ExitIfError(err)
	_ = cmd.Wait()
}
