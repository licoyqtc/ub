package lightdev

import (
	"github.com/cybergarage/go-net-upnp/net/upnp/http"
	"fmt"
	"ubeybox/api"
)

type pHandler func(*http.Request,http.ResponseWriter)
type Router map[string]pHandler



type HttpHandler struct {
	httpRouter map[string]Router
}

func NewHttpHandler() *HttpHandler{
	pHttpHandler := &HttpHandler{}
	pHttpHandler.httpRouter = make(map[string]Router)

	pHttpHandler.RouterGet("/register" , api.Register)
	pHttpHandler.RouterGet("/login" , api.Login)
	pHttpHandler.RouterGet("/diskinfo" , api.Getdiskinfo)
	pHttpHandler.RouterGet("/testsql" , api.TestSql)

	return pHttpHandler
}


func (ph *HttpHandler) RouterGet(path string , handler pHandler){
	if len(ph.httpRouter[http.GET]) == 0 {
		rMap := make(Router)
		rMap[path] = handler
		ph.httpRouter[http.GET] = rMap
	} else {
		ph.httpRouter[http.GET][path] = handler
	}

}

func (ph *HttpHandler) RouterPost(path string , handler pHandler){
	if len(ph.httpRouter[http.POST]) == 0 {
		rMap := make(Router)
		rMap[path] = handler
		ph.httpRouter[http.POST] = rMap
	} else {
		ph.httpRouter[http.POST][path] = handler
	}

}


func (h HttpHandler) HTTPRequestReceived(req *http.Request,rsp http.ResponseWriter) {
	fmt.Printf("get http req : url %+v body : %+v\n",*((*req).URL) , (*req).Body )

	f := h.httpRouter[req.Method][req.URL.Path]
	if f == nil {
		rsp.WriteHeader(http.StatusNotFound)
		return
	}

	f(req,rsp)

}