package main

import (
	"math"
)

type Order struct {
	Qty  float64 `json:"quantity"`
	Rate float64 `json:"rate"`
}

type Exchange struct {
	Asks []Order
	Bids []Order
}

type Cart struct {
	Qty    float64 `json:"quantity"`
	Amount float64 `json:"amount"`
}

func NewExchange() *Exchange {
	return &Exchange{Bids: Bids, Asks: Asks}
}

func (e *Exchange) addBid(order Order) {
	idx := len(e.Bids)
	for id, ord := range e.Bids {
		if order.Rate > ord.Rate {
			idx = id
			break
		}
	}

	e.Bids = append(e.Bids[:idx], append([]Order{order}, e.Bids[idx:]...)...)
}

func (e *Exchange) addAsk(order Order) {
	idx := len(e.Asks)
	for id, ord := range e.Asks {
		if order.Rate < ord.Rate {
			idx = id
			break
		}
	}

	e.Asks = append(e.Asks[:idx], append([]Order{order}, e.Asks[idx:]...)...)
}

func executeOrder(orders []Order, qty float64) (Cart, []Order) {
	cart := Cart{}
	flagId := 0
	for id, ord := range orders {

		if qty == 0 {
			break
		}

		matchQty := math.Min(qty, ord.Qty)

		orders[id].Qty -= matchQty
		cart.Qty += matchQty
		cart.Amount += (matchQty * ord.Rate)
		qty -= matchQty

		if orders[id].Qty == 0 {
			flagId = id + 1
		}

	}
	return cart, orders[flagId:]
}

func (e *Exchange) marketOrder(orderType string, qty float64) Cart {
	switch orderType {
	case "BUY":
		cart, orders := executeOrder(e.Asks, qty)
		e.Asks = orders
		return cart

	case "SELL":
		cart, orders := executeOrder(e.Bids, qty)
		e.Bids = orders
		return cart

	default:
		return Cart{}
	}
}

func (e *Exchange) limitOrder(orderType string, order Order) Cart {
	if orderType == "BUY" {
		cart := Cart{}
		flagId := 0

		for id, exOrder := range e.Asks {
			if order.Qty == 0 {
				break
			}

			if order.Rate < exOrder.Rate {
				e.addBid(order)
				break
			}

			matchQty := math.Min(order.Qty, exOrder.Qty)
			cart.Qty += matchQty
			cart.Amount += (matchQty * exOrder.Rate)
			e.Asks[id].Qty -= matchQty
			order.Qty -= matchQty

			if e.Asks[id].Qty == 0 {
				flagId = id + 1
			}

		}

		e.Asks = e.Asks[flagId:]
		return cart
	} else if orderType == "SELL" {

		cart := Cart{}
		flagId := 0

		for id, exOrder := range e.Bids {

			if order.Qty == 0 {
				break
			}

			if order.Rate > exOrder.Rate {
				e.addAsk(order)
				break
			}

			matchQty := math.Min(order.Qty, exOrder.Qty)
			cart.Qty += matchQty
			cart.Amount += (matchQty * exOrder.Rate)
			e.Bids[id].Qty -= matchQty
			order.Qty -= matchQty

			if e.Bids[id].Qty == 0 {
				flagId = id + 1
			}

		}
		e.Bids = e.Bids[flagId:]
		return cart

	} else {
		return Cart{}
	}
}

func (e *Exchange) getBids() []Order {
	return e.Bids
}

func (e *Exchange) getAsks() []Order {
	return e.Asks
}
