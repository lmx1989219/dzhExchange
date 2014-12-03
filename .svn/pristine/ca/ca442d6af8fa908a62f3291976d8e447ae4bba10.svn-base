package exchange

import (
	"atmmessage"
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"github.com/pebbe/zmq4"
//	"time"
	"strings"
)

//订阅行情
//行情推送来触发撮合线程运行
func WakeUp(codes []string) {
	ctx, _ := zmq4.NewContext()
	socket, _ := ctx.NewSocket(zmq4.SUB)
	defer socket.Close()
	socket.Connect("tcp://10.15.89.122:6666")
	socket.SetSubscribe("")
	fmt.Println("quote sub succed")
	for {
		resp, _ := socket.RecvBytes(0)
		if len(resp) != 0 {
			tobj := &atmmessage.Did60130{}
			//unswap obj
			e := proto.Unmarshal(resp, tobj)
			if nil == e {
				//fmt.Println("quote data is:", tobj)
				for i := range codes {
					if strings.EqualFold(tobj.GetObj(), codes[i]) {
						quote := new(Quote)
						quote.Code = tobj.GetObj()
						quote.Name = tobj.GetName()
						mpool := MatchPool
						queue := mpool[quote.Code]
						//开启一个go程去处理
						go Ss.DealMatch(queue, quote, DealQueue)
					}

				}

			} else {
				//log.Fatalln("decode fail ", e)
			}

		}
//		time.Sleep(1 * time.Millisecond)
	}
}
