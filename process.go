/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-15 12:10
**/

package main

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/text"
)

type Child struct {
	Pid     int
	Restart bool
	Status  string
}

type Proc struct {
	Name     string
	Children []*Child
	Cmd      []string
	OutPath  string
	ErrPath  string

	ch chan struct{}
}

type Process []*Proc

func (p Process) String() string {

	if len(p) == 0 {
		return ""
	}

	var nameMaxLen = 0
	var cmdMaxLen = 0
	var pidMaxLen = 0
	var statusMaxLen = 0

	for i := 0; i < len(p); i++ {
		var nl = text.RuneCount(p[i].Name)
		if nl > nameMaxLen {
			nameMaxLen = nl
		}

		for j := 0; j < len(p[i].Cmd); j++ {
			var cl = text.RuneCount(p[i].Cmd[j])
			if cl > cmdMaxLen {
				cmdMaxLen = cl
			}

			var pl = text.RuneCount(fmt.Sprintf("%d", p[i].Children[j].Pid))
			if pl > pidMaxLen {
				pidMaxLen = pl
			}

			var sl = text.RuneCount(p[i].Children[j].Status)
			if sl > statusMaxLen {
				statusMaxLen = sl
			}
		}
	}

	nameMaxLen += 2
	cmdMaxLen += 2
	pidMaxLen += 2
	statusMaxLen += 2

	var str = "\n"

	for i := 0; i < len(p); i++ {

		str += p[i].Name + strings.Repeat(" ", nameMaxLen-text.RuneCount(p[i].Name))

		for j := 0; j < len(p[i].Cmd); j++ {
			if j == 0 {
				str += p[i].Cmd[j]
			} else {
				str += strings.Repeat(" ", nameMaxLen) + p[i].Cmd[j]
			}

			str += strings.Repeat(" ", cmdMaxLen-text.RuneCount(p[i].Cmd[j])) +
				fmt.Sprintf("%d", p[i].Children[j].Pid) +
				strings.Repeat(" ", pidMaxLen-text.RuneCount(fmt.Sprintf("%d", p[i].Children[j].Pid))) +
				p[i].Children[j].Status +
				"\n"
		}

		if i != len(p)-1 {
			str += "\n"
		}
	}

	return str
}
