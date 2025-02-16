package main

import(
	"fmt"
	"net"
	"bufio"
	"os"
)

func main(){
	fmt.Println("客户端启动...")
	conn,err := net.Dial("tcp","192.168.0.105:8888")
	if err != nil{
		fmt.Println("连接失败,err:",err)
		return
	}
	fmt.Println("连接成功, conn:",conn)
	
	reader := bufio.NewReader(os.Stdin)

	str,err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("读取失败,err:",err)
	}

	n,err := conn.Write([]byte(str))
	if err != nil{
		fmt.Println("连接失败,err:",err)
	}
	fmt.Printf("发送成功，发送了%d字节数据\n",n)
}


