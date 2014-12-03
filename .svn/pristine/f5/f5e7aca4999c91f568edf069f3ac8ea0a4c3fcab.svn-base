package exchange

type StkMatchQueue struct {
	Queue MatchQueue
}

func (stm *StkMatchQueue) InitQueue() {
	stm.Queue.BuyQueue = make(map[*CommonKey]*CommonOrder)
	stm.Queue.SellQueue = make(map[*CommonKey]*CommonOrder)
}

func (stm *StkMatchQueue) InsertOrder(o *CommonOrder) {
	ckey := new(CommonKey)
	ckey.Oid = o.Oid
	ckey.Price = o.Price
	if o.BsFlag == 0 {
		stm.Queue.BuyQueue[ckey] = o
	} else {
		stm.Queue.SellQueue[ckey] = o
	}
}

func (stm *StkMatchQueue) delOrder(o CommonKey) {

}
