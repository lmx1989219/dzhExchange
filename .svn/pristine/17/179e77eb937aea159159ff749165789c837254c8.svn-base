package exchange

type CommonOrder struct {
	Oid    int     //订单id
	Price  float32 //订单价格
	Code   string  //合约代码
	Name   string  //合约名称
	BsFlag int     //买卖标记
	Time   string  //下单时间
}

type Quote struct {
	Code string    //合约代码
	Name string    //合约名称
	Buy  []float32 //买档
	Sell []float32 //卖档

	NewPrice    []float32 //最新价
	LastClose   []float32 //昨收价
	SettlePrice []float32 //收盘价

	Date string //行情时间
}

type MatchQueue struct {
	BuyQueue  map[*CommonKey]*CommonOrder //should be implement sorted map desc
	SellQueue map[*CommonKey]*CommonOrder //should be implement sorted map asc
}

type CommonKey struct {
	Oid   int
	Price float32
}

type Contracts struct{
	Code string    //合约代码
	Name string    //合约名称
}