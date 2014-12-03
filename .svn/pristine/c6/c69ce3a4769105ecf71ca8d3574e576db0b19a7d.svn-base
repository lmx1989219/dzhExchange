package exchange

//撮合接口
var Ss = new(StkMatchService)

//处理请求信号
var Sign = make(chan int, 10*1024*1024)

//处理deal信号
var DealQueue = make(chan *CommonOrder, 10*1024*1024)


//处理行情信号
var SignQuote = make(chan int, 2)

//撮合池
var MatchPool map[string]*StkMatchQueue

var Stqueue *StkMatchQueue

//code table
var CodesTable []string