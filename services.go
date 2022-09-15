/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-15 03:17
**/

package main

import (
	"fmt"
)

func findServiceByName(name string) *Proc {
	var list = getService()

	var service *Proc

	for i := 0; i < len(list); i++ {
		if list[i].Name == name {
			service = list[i]
			break
		}
	}

	if service == nil {
		return nil
	}
	return service
}

func services() {
	switch Args(2) {
	case "list":
		fmt.Println(getService())
	case "log":
		logService()
	case "err":
		errService()
	case "stop":
		stopService()
	case "stopAll":
		stopAllServices()
	case "start":
		startService()
	case "restart":
		restartService()
	case "":
		fmt.Println(getService())
	case "-h", "--help":
		fmt.Println(helpService())
	default:
		fmt.Println(Help())
	}
}

func stopAllServices() {
	var bts, err = httpGet("stopAll", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bts))
}

func restartService() {
	var name = GetArgs([]string{"restart"})
	if name == "" {
		fmt.Println(Help())
		return
	}

	var bts, err = httpGet("restart", M{"name": name})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bts))
}

func startService() {
	var name = GetArgs([]string{"start"})
	if name == "" {
		fmt.Println(Help())
		return
	}

	var bts, err = httpGet("start", M{"name": name})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bts))
}

func stopService() {
	var name = GetArgs([]string{"stop"})
	if name == "" {
		fmt.Println(Help())
		return
	}

	var bts, err = httpGet("stop", M{"name": name})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bts))
}
