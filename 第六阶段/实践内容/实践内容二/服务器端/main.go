package main
import(
	"fmt"
	"net"
)

func process(conn net.Conn){
	//关闭连接
	defer conn.Close()
	for{//创建切片用于存储客户端发送的数据
		buf := make([]byte,1024)
		n,err := conn.Read(buf)//Read函数用于读取客户端发送的数据
		if err != nil{
			return
		}
		//将读取到的数据转换为字符串并输出
		//buf[:n]用于截取buf切片中的前n个元素
		fmt.Println(string(buf[:n]))
	}
}

func main(){
	fmt.Println("服务器端启动..")
	//监听，需要指定协议和IP地址和端口号
	listen,err := net.Listen("tcp","192.168.0.105:8888")
	if err != nil{//监听失败
		fmt.Println("监听失败, err:",err)
		return
	}
	//监听成功，等待客户端连接
	// Accept()函数会阻塞等待客户端连接，有客户端连接后返回连接对象
	for{
		conn,err2 := listen.Accept()
		if err2 != nil{
			fmt.Println("客户端等待失败,err2:",err2)
		}else{//连接成功
			//RemoteAddr()用于获取客户端的IP地址和端口号
			fmt.Printf("等待连接成功,con=%v , 接收到的客户端: %v \n",conn,conn.RemoteAddr().String())
		}
        //协程处理客户端请求
		go process(conn)
	}
}