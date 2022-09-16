/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-16 23:22
**/

package program

import (
	"fmt"

	"github.com/kardianos/service"
	"github.com/lemonyxk/pm/def"
	"github.com/lemonyxk/pm/tools"
)

type Program struct{}

func (p *Program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *Program) run() {
	// Do work here
	run()
}

func (p *Program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	_, err := tools.HttpGet(def.EXIT, nil)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
