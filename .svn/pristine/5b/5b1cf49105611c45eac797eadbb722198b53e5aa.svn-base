package exchange

type MatchService interface {
	//撮合逻辑处理
	DealMatch(queue *StkMatchQueue,q *Quote,f chan *CommonOrder)
}

type QueueService interface {
	//初始化队列
	InitQueue()
	//订单入队列
	InsertOrder(o *CommonOrder)
	//删除某个key
	DelOrder(o *CommonKey)
}
