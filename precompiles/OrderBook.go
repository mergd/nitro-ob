// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package precompiles

import (
	ob "github.com/muzykantov/orderbook"
)

type OrderBook struct {
	addresss addr // 0x63
	ob       *ob.OrderBook
}

// func (ob *OrderBook) ProcessLimitOrder(side Side, orderID string, quantity, price decimal.Decimal) (done []*Order, partial *Order, err error) { ... }

// func (ob *OrderBook) ProcessMarketOrder(side Side, quantity decimal.Decimal) (done []*Order, partial *Order, quantityLeft decimal.Decimal, err error) { .. }

// func (ob *OrderBook) CancelOrder(orderID string) *Order { ... }
// func (ob *OrderBook) Order(orderID string) *Order {
// func (ob *OrderBook) Depth() (asks, bids []*PriceLevel) {
// func (ob *OrderBook) CalculateMarketPrice(side Side, quantity decimal.Decimal) (price decimal.Decimal,quant decimal.Decimal, err error) {

// Some functions are only viewable view RPC methods
// func (ob *OrderBook) GetOrderSide(side Side) *OrderSide {

// func (ob *OrderBook) MarketOverview() *MarketView {

// func compileOrders(orders *OrderSide) map[string]decimal.Decimal {
