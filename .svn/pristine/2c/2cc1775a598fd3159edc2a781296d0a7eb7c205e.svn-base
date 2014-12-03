package exchange

import (
	"atmmessage"
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"github.com/pebbe/zmq4"
	"log"
	"time"
)

func PushBack() {
	//    required int32 code = 1;
	//    required double dealPrice = 2;
	//    required int32 dealCount = 3;
	//    required int32 marketCode = 4;
	//    required string userName = 5;
	//    required int32 orderId = 6;
	//	for{
	//		<-DealQueue
	//	}
	ctx, _ := zmq4.NewContext()
	publisher, _ := ctx.NewSocket(zmq4.PUB)
	defer publisher.Close()
	publisher.Bind("tcp://*:6668")
	fmt.Println("server is start...")

	for v := range DealQueue {
		if nil != v {
			user := &atmmessage.M2C{
				Code:       proto.Int32(6000),
				UserName:   proto.String("test"),
				DealPrice:  proto.Float64(5.25),
				DealCount:  proto.Int32(6000),
				MarketCode: proto.Int32(1),
				OrderId:    proto.Int32(1),
			}
			//swap obj
			encObj, err := proto.Marshal(user)
			if nil == err {
				publisher.SendBytes(encObj, 0)
				fmt.Println("send to counter succed..")
				time.Sleep(1 * time.Millisecond)
			} else {
				log.Fatalln("encode fail", err)

			}
		}
	}
}
