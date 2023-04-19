/**
* @program: pm
*
* @description:
*
* @author: lemon
*
* @create: 2022-09-12 16:13
**/

package main

import (
	"fmt"

	"github.com/lemonyxk/pm/def"
	"github.com/lemonyxk/pm/tools"
)

func main() {
	switch tools.Args(1) {
	case def.EMPTY, def.LIST:
		list()
	case def.LOG:
		logService()
	case def.ERR:
		errService()
	case def.STOP:
		fmt.Println(Query(def.STOP))
	case def.SHUTDOWN:
		fmt.Println(Query(def.SHUTDOWN))
	case def.START:
		fmt.Println(Query(def.START))
	case def.RESTART:
		fmt.Println(Query(def.RESTART))
	case def.REMOVE:
		fmt.Println(Query(def.REMOVE))
	case "-h", "--help":
		fmt.Println(tools.PMHelp())
	default:
		fmt.Println(tools.PMHelp())
	}
}
