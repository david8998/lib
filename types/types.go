package types

import "encoding/json"

type OrderBook struct {
	Symbol string      `json:"symbol"`
	Time   int64       `json:"time"`
	Ask    [][2]string `json:"ask"` //价格，数量 卖盘
	Bid    [][2]string `json:"bid"` //价格，数量 买盘
}

func (o OrderBook) MarshalBinary() (data []byte, err error) {
	return json.Marshal(o)
}
