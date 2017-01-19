//自定义中间件 recorder记录
package main

import (
    "net/http"
    "net/http/httptest"
)

type MiddleWare struct {
    http.Handler
}

func (self *MiddleWare)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
    rec := httptest.NewRecorder()
    self.Handler.ServeHTTP(rec,r)

    for k,v := range rec.Header(){
        w.Header()[k] = v
    }
    w.Header().Set("go-web-foundation","vip")
    w.WriteHeader(418)
    w.Write([]byte("hey!"))
    w.Write(rec.Body.Bytes())
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func main() {
    mid := &MiddleWare{
        http.HandlerFunc(myHandler),
    }
    http.ListenAndServe(":8080",mid)
}