/**
* @program: pm
*
* @description:
*
* @author: lemon
*
* @create: 2022-09-12 16:19
**/

package tools

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/lemonyxk/console"
	"github.com/lemonyxk/pm/config"
	"github.com/lemonyxk/utils/v3"
	"github.com/shirou/gopsutil/v3/process"
)

type M map[string]string

func Args(index uint) string {
	if len(os.Args) < int(index)+1 {
		return ""
	}
	return os.Args[index]
}

func HasArgs(flag string) bool {
	var args = os.Args
	for i := 0; i < len(args); i++ {
		if args[i] == flag {
			return true
		}
	}
	return false
}

func GetArgs(flag []string) string {
	var args = os.Args
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(flag); j++ {
			if args[i] == flag[j] {
				if i+1 < len(args) {
					return args[i+1]
				}
			}
		}
	}
	return ""
}

func FindProcess(pid ...int32) []*process.Process {

	if len(pid) == 0 {
		return nil
	}

	var ps, err = process.Processes()
	if err != nil {
		console.Exit(err)
	}

	var res []*process.Process
	for i := 0; i < len(ps); i++ {
		var p = ps[i]
		if utils.ComparableArray(&pid).Has(p.Pid) {
			res = append(res, p)
			if len(pid) == len(res) {
				return res
			}
		}
	}

	return res
}

func GetConfig(name string) config.Config {
	for i := 0; i < len(config.Configs); i++ {
		if config.Configs[i].Name == name {
			return config.Configs[i]
		}
	}
	return config.Config{}
}

func HttpGet(path string, params map[string]string) ([]byte, error) {
	Url, err := url.Parse("http://127.0.0.1:52525/" + path)
	if err != nil {
		return nil, err
	}

	var p = url.Values{}
	for k, v := range params {
		p.Set(k, v)
	}

	var pStr = p.Encode()

	Url.RawQuery = pStr

	res, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		return nil, err
	}

	req, err := http.DefaultClient.Do(res)
	if err != nil {
		return nil, err
	}

	defer func() { _ = req.Body.Close() }()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func ExitIfError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func FixURL(p string) string {
	if len(p) == 0 {
		return ""
	}
	return p[1:]
}
