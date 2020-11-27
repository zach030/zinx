package main

import (
	"Zinx/ziface"
	"Zinx/znet"
	"fmt"
)

//ping test
type PingRouter struct {
	znet.BaseRouter
}

// test handle router
func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call router: handle")
	//先读取客户端数据，再回写
	fmt.Println("recv client msg id = ", request.GetMsgID(),
		" data = ", string(request.GetData()))
	err := request.GetConnection().SendMsg(1,[]byte("ping"))
	if err != nil {
		fmt.Println("call back  ping error:", err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}
// test handle router
func (p *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call router: Hello zinx router handle")
	//先读取客户端数据，再回写
	fmt.Println("recv client msg id = ", request.GetMsgID(),
		" data = ", string(request.GetData()))
	err := request.GetConnection().SendMsg(1,[]byte("Hello zinx"))
	if err != nil {
		fmt.Println("call back  ping error:", err)
	}
}

func DoConnBegin(connection ziface.IConnection){
	fmt.Println("==========> Do connection Begin")
	if err := connection.SendMsg(202,[]byte("Do connection Begin"));err != nil{
		fmt.Println("Call hook func failed:",err)
		return
	}
}

func DoConnStop(connection ziface.IConnection){
	fmt.Println("==========> Do connection Stop, conn id = ",connection.GetConnID())
}

/*基于zinx框架开发的服务器应用程序*/
func main() {
	// 1 创建server句柄
	s := znet.NewServer("zinx V0.9")
	// 2 添加自定义router
	s.AddRouter(0,&PingRouter{})
	s.AddRouter(1,&HelloZinxRouter{})
	// 3 注册hook回调func
	s.SetOnConnStart(DoConnBegin)
	s.SetOnConnStop(DoConnStop)
	// 4 启动server
	s.Serve()
}
