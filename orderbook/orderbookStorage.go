package orderbook

import (
	"github.com/ethereum/go-ethereum/common"
	ob "github.com/muzykantov/orderbook"
)

type OrderbookStorage struct {
	OrderbookWrappers []*OrderbookWrapper
}

type Token struct {
	Address  common.Address
	Symbol   string
	Decimals uint8
}

type OrderbookWrapper struct {
	Orderbook *ob.OrderBook
	Token0    Token
	Token1    Token
	Scalar    uint16
}

func NewOrderbook(token0, token1 Token, scalar uint16) *OrderbookWrapper {
	return &OrderbookWrapper{
		Orderbook: &ob.OrderBook{},
		Token0:    token0,
		Token1:    token1,
		Scalar:    scalar,
	}
}
