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
	case "remove":
		removeService()
	case "active":
		activeService()
	case "unActive":
		unActiveService()
	case "load":
		loadService()
	case "":
		fmt.Println(getService())
	case "-h", "--help":
		fmt.Println(helpService())
	default:
		fmt.Println(helpService())
	}
}

func loadService() {
	var bts, err = httpGet("load", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bts))
}

func unActiveService() {
	var name = GetArgs([]string{"unActive"})
	var bts, err = httpGet("unActive", M{"name": name})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bts))
}

func activeService() {
	var name = GetArgs([]string{"active"})
	var bts, err = httpGet("active", M{"name": name})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bts))
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
	var bts, err = httpGet("restart", M{"name": name})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bts))
}

func removeService() {
	var name = GetArgs([]string{"remove"})
	var bts, err = httpGet("remove", M{"name": name})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bts))
}

func startService() {
	var name = GetArgs([]string{"start"})
	var bts, err = httpGet("start", M{"name": name})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bts))
}

func stopService() {
	var name = GetArgs([]string{"stop"})
	var bts, err = httpGet("stop", M{"name": name})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bts))
}
