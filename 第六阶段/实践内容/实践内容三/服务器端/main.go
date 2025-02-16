package main

import(
	"net/http"
	"log"
)//log用于记录错误信息，net/http用于构建 HTTP 服务器


 //处理客户端请求  
 //w http.ResponseWriter用于向客户端发送响应
 //r *http.Request用于获取客户端请求的信息
func handler (w http.ResponseWriter, r *http.Request) {  
    w.Write([]byte("Hello, World!")) 
}  

func main() {  
    //将handler函数注册到根路径/上 
    http.HandleFunc("/", handler) 
	
    log.Println("Server starting on port 8080...")  
    err := http.ListenAndServe(":8080", nil)  //启动服务器
    if err != nil {  
        log.Fatal("ListenAndServe: ", err)  
    }  
}  