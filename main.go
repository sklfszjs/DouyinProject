// Code generated by hertz generator.

package main

import (
	"fmt"

	utils "github.com/cloudwego/biz/utils"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func init() {
	utils.Form_tables()
}

func main() {
	// go utils.RunFileServer()
	h := server.New(server.WithMaxRequestBodySize(400*1024*1024),
		server.WithHostPorts(fmt.Sprintf("%s:%d", utils.GetConfigs().IP, utils.GetConfigs().Port)))
	register(h)
	h.Spin()
}
