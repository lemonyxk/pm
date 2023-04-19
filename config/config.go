/**
* @program: pm
*
* @description:
*
* @author: lemon
*
* @create: 2022-09-16 23:13
**/

package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/lemonyxk/console"
	"github.com/lemonyxk/utils/v3"
)

var Configs []Config

type Config struct {
	Name    string   `json:"name"`
	User    string   `json:"user"`
	Dir     string   `json:"dir"`
	Command []string `json:"command"`
	Restart bool     `json:"restart"`
	Out     string   `json:"out"`
	Err     string   `json:"err"`
}

func InitConfig() []Config {

	Configs = []Config{}

	files, err := os.ReadDir(CfgDir)
	if err != nil {
		console.Exit(err)
	}

	for i := 0; i < len(files); i++ {
		var fullPath = filepath.Join(CfgDir, files[i].Name())
		if files[i].IsDir() {
			continue
		}

		if !strings.HasSuffix(fullPath, ".json") {
			continue
		}

		var f = utils.File.ReadFromPath(fullPath).Bytes()

		var c Config
		err = utils.Json.Decode(f, &c)
		if err != nil {
			console.Info(err)
			continue
		}

		var n = files[i].Name()

		c.Name = n[:len(n)-len(filepath.Ext(n))]

		Configs = append(Configs, c)
	}

	return Configs
}
