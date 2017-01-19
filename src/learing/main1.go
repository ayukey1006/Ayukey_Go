// //限制请求来源中间件
// package main

// import (
// 	"net/http"
// )

// type SingleHost struct {
// 	handler     http.Handler
// 	allowedHost string
// }

// func (self *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if r.Host == self.allowedHost {
// 		self.handler.ServeHTTP(w, r)
// 	} else {
// 		w.WriteHeader(403)
// 	}
// }

// func myHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello world"))
// }

// func main() {
// 	single := &SingleHost{
// 		handler:     http.HandlerFunc(myHandler),
// 		allowedHost: "localhost:8080",
// 	}

// 	http.ListenAndServe(":8080", single)
// }
