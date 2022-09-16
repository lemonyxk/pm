//go:build darwin
// +build darwin

/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-12 18:24
**/

package config

func HomePath() string {
	return "/Library/Application Support"
}
