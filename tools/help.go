/**
* @program: pm
*
* @description:
*
* @author: lemon
*
* @create: 2022-09-12 16:17
**/

package tools

import (
	"github.com/jedib0t/go-pretty/table"
	"github.com/lemonyxk/console"
)

var options = table.Options{
	DrawBorder:      false,
	SeparateColumns: false,
	SeparateFooter:  false,
	SeparateHeader:  false,
	SeparateRows:    false,
}

func PMDHelp() string {
	var t = console.NewTable()
	t.Style().Options = options
	t.Row("pmd install", "install pmd to launch")
	t.Row("pmd uninstall", "uninstall pmd")
	t.Row("pmd start", "start pmd server")
	t.Row("pmd stop", "stop pmd server")
	t.Row("pmd restart", "restart pmd server")
	t.Row("pmd status", "get pmd server status")
	t.Row("pmd info", "get pmd server info")
	t.Row("pmd log", "get pmd log")
	t.Row("pmd err", "get pmd err")
	return t.Render()
}

func PMHelp() string {
	var t = console.NewTable()
	t.Style().Options = options
	t.Row("pm list", "list services")
	t.Row("pm log [service name]", "service log")
	t.Row("pm err [service name]", "service err")
	t.Row("pm stop [service name]", "stop service")
	t.Row("pm shutdown [service name]", "stop all service")
	t.Row("pm start [service name]", "start service")
	t.Row("pm restart [service name]", "restart service")
	t.Row("pm remove [service name]", "remove service")
	return t.Render()
}
