/**
* @program: pm
*
* @description:
*
* @author: lemon
*
* @create: 2023-04-19 22:39
**/

package app

import (
	"github.com/kardianos/service"
	"github.com/lemonyxk/pm/process"
	hash "github.com/lemonyxk/structure/map"
)

var SigMap = hash.NewSync[string, *process.Proc]()
var Server service.Service
