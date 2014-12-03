package exchange

import (
//	"time"
//	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"github.com/pebbe/zmq4"
	"math/rand" 
	
)

//请求处理线程
func StartProcess(sign chan int) {
	ctx, _ := zmq4.NewContext()
	socket, _ := ctx.NewSocket(zmq4.SUB)
	defer socket.Close()
	socket.Connect("tcp://127.0.0.1:6665")
	//WT_ORDER_REQ_QUEUE
	socket.SetSubscribe("")
	fmt.Println("request order sub succed")
	codesArr := CodesTable
	for {
		socket.RecvBytes(0)
//		fmt.Println("cur quest",string(resp))
		co := new(CommonOrder)
		co.Oid = 10000
		co.Code = codesArr[rand.Int63n(2000)+1]
		co.Price = 8.2
		co.BsFlag = 0
		stQueue := MatchPool[co.Code]
		stQueue.InsertOrder(co)
		//写入
		sign <- 1
//		time.Sleep(1 * time.Millisecond)
	}
}
