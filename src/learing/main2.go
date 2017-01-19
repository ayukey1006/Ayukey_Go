// //请求内容追加中间件
// package main

// import (
// 	"net/http"
// )

// type MiddleWare struct {
// 	http.Handler
// }

// func (self *MiddleWare)ServeHTTP(w http.ResponseWriter, r *http.Request)  {
//     self.Handler.ServeHTTP(w,r)
//     w.Write([]byte("JC"))
// }

// func myHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello world!"))
// }

// func main() {
//     mid := &MiddleWare{
//         http.HandlerFunc(myHandler),
//     }
//     http.ListenAndServe(":8080",mid)
// }
