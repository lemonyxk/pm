/**
* @program: pm
*
* @description:
*
* @author: lemo
*
* @create: 2022-09-16 20:18
**/

package program

import (
	"fmt"
	"net"
	"net/http"

	"github.com/lemonyxk/console"
	"github.com/lemonyxk/utils/v3"
)

func (h *handler) end(w http.ResponseWriter, v any) {
	w.WriteHeader(200)
	_, _ = w.Write(utils.Json.Encode(v))
}

func (h *handler) endStr(w http.ResponseWriter, v any) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte(fmt.Sprintf("%v", v)))
}

func CreateServer() error {
	var err error
	var netListen net.Listener
	var server = http.Server{Addr: ":52525", Handler: &handler{}}

	netListen, err = net.Listen("tcp", server.Addr)

	if err != nil {
		return err
	}

	go func() {
		err = server.Serve(netListen)
		if err != nil {
			console.Info(err)
		}
	}()

	return nil
}
