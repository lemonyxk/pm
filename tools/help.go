/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-12 16:17
**/

package tools

func PMDHelp() string {
	return `
pmd install
  -- install pmd into launch

pmd uninstall
  -- uninstall pmd from launch

pmd start
  -- start pmd server

pmd stop
  -- stop pmd server

pmd restart
  -- restart pmd server

pmd status
  -- get pmd server status

pmd run
  -- run pmd server

pmd info
  -- get pmd server info

pmd log
  -- get pmd log

pmd err
  -- get pmd err

pmd err
  -- get pmd err`
}

func PMHelp() string {
	return `
pm
  -- list services

pm list
  -- list services

pm log [service name]
  -- service log

pm err [service name]
  -- service err

pm stop [service name]
  -- stop service

pm stopAll
  -- stop all service

pm start [service name]
  -- start service

pm restart [service name]
  -- restart service

pm remove [service name]
  -- remove service

pm active [service name]
  -- active service

pm unAction [service name]
  -- unAction service`
}
