package main

import(
	"fmt"
	"net"
	"bufio"
	"os"
)

func main(){
	fmt.Println("客户端启动...")
	//调用Dial函数，参数需指定tcp协议，服务器IP地址和端口号
	//Dial函数返回两个值，一个是连接对象，一个是错误信息
	conn,err := net.Dial("tcp","192.168.0.105:8888")
	if err != nil{
		fmt.Println("连接失败,err:",err)
		return
	}
	fmt.Println("连接成功, conn:",conn)
	//通过客户端发送数据
	reader := bufio.NewReader(os.Stdin)//os.Stdin代表终端标准输入
    //读取终端输入的内容
	str,err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("读取失败,err:",err)
	}
    //将数据发送给服务器
	//Write函数用于发送数据，参数为字节切片，返回两个值，一个是发送的字节数，一个是错误信息
	n,err := conn.Write([]byte(str))
	if err != nil{
		fmt.Println("连接失败,err:",err)
	}
	fmt.Printf("发送成功，发送了%d字节数据\n",n)
}


