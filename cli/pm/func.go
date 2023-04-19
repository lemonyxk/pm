/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2023-04-19 23:22
**/

package main

import (
	"fmt"

	"github.com/lemonyxk/console"
)

func list() {
	var list = GetServices()
	var table = console.NewTable()
	table.Style().Options.DrawBorder = false
	// table.Header("name", "cmd", "pid", "status", "time")
	for i := 0; i < len(list); i++ {
		var process = list[i]
		for j := 0; j < len(process.Children); j++ {
			var child = process.Children[j]
			var name = process.Name
			var status = child.Status
			if j != 0 {
				name = ""
			}
			if child.Status == "running" {
				status = console.FgGreen.Sprint(child.Status)
			} else {
				status = console.FgRed.Sprint(child.Status)
			}
			table.Row(name, process.Cmd[j], child.Pid, status, child.Time.Format("01-02 15:04:05"))
		}
		if i != len(list)-1 {
			table.Row("", "", "", "", "")
		}
	}
	fmt.Println(table.Render())
}
