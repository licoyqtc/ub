package api

import (
	"github.com/cybergarage/go-net-upnp/net/upnp/http"
	"fmt"
)

func Register(req *http.Request,rsp http.ResponseWriter) {


	rsp.Write([]byte(fmt.Sprintf("receive get req done :%s",req.URL.Path)))
}
