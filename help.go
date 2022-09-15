/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-12 16:17
**/

package main

func Help() string {
	return `
pm install
  -- install pm into launch

pm uninstall
  -- uninstall pm from launch

pm start
  -- start pm server

pm stop
  -- stop pm server

pm restart
  -- restart pm server

pm status
  -- get pm server status

pm run
  -- run pm server

pm info
  -- get pm server info

pm log
  -- get pm log

pm err
  -- get pm err

pm err
  -- get pm err` + "\n" + helpService()
}

func helpService() string {
	return `
pm services
  -- list services

pm service list
  -- list services

pm service log [service name]
  -- service log

pm service err [service name]
  -- service err

pm service stop [service name]
  -- stop service

pm service stopAll
  -- stop all service

pm service start [service name]
  -- start service

pm service restart [service name]
  -- restart service`
}
