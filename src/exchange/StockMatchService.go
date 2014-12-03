package exchange

import (
	"fmt"
	"sync"
)

type FuMatchService struct {
	id    string
	isRun bool
}

type StkMatchService struct {
	id    string
	isRun bool
}

//single write ,more read  slock
var l sync.RWMutex

func (stm *StkMatchService) DealMatch(queue *StkMatchQueue, q *Quote, f chan *CommonOrder) {
	//lock queue util deal end
	l.Lock()
	buy := queue.Queue.BuyQueue
	sell := queue.Queue.SellQueue
	var dealOrder *CommonOrder = new(CommonOrder)
	for _, v := range buy {
		fmt.Println("stk order didn't match", v)
	}
	for _, v := range sell {
		fmt.Println("stk order didn't match", v)
	}
	defer l.Unlock()
	f <- dealOrder
}
