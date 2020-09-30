package main

import(
    "fmt"
    "net/http"
    )
func hello(w http.ResponseWriter,req *http.Request){
    //handler,using http.HandlerFunc with signature
    
    fmt.Fprintf(w,"hello\n")//Fprintf往第一个参数写入字符串，第一个参数实现io interface
    //responsewtriter用这个字符串fill in the response,把请求头回响到response body
}

func headers(w http.ResponseWriter,req *http.Request){
    for name,header := range req.Header{
        for _,h := range header{
            fmt.Fprintf(w,"%v:%v\n",name,h)
        }
    }
}

func main() {
    http.HandleFunc("/", hello)
    //HandlerFun:在服务器路由注册handler，设置为库里面的默认路由，并接受一个handler
    //http.HandleFunc("/", headers)
    http.ListenAndServe(":8080", nil)
}

