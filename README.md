Agenda: Build Exchange from scratch

Constrain: Currently only single currency support.

Asks: 
Bids:

Matching Engine:
Order types: Market, Limit



<!--
// #TODO: complete this
func (e *Exchange) limitOrder(orderType string, order Order) {
	if orderType == "BUY" {
		purchase := Order{}
		_ = purchase

		for _, ord := range e.Bids {
			if ord.amount > order.amount {
				break
			}
		}
	} else if orderType == "SELL" {
	} else {
		return
	}
}
-->
