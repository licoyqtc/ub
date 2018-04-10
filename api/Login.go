package api

import (
	"github.com/cybergarage/go-net-upnp/net/upnp/http"
	"fmt"

	"encoding/json"
)


type lgRsp struct {
	Err_no  int		`json:"err_no"`
	Err_msg string	`json:"err_msg"`
}

func Login(req *http.Request,rsp http.ResponseWriter) {

	ret := lgRsp{}
	ret.Err_msg = fmt.Sprintf("receive post req done :%s",req.URL.Path)

	r , _ := json.Marshal(ret)
	rsp.Write(r)
}
