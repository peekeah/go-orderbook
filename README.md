# Go Orderbook

Central limit order book (CLOB) engine in Go, supporting market and limit orders.

---

## Features

- Sorted order book (bids descending, asks ascending)
- Market order execution with partial fills
- Limit order matching with residual order placement

---

## Order Types

| Type   | Behaviour                                                                 |
|--------|---------------------------------------------------------------------------|
| Market | Executes immediately against best available prices; partial fills allowed |
| Limit  | Matches up to the specified rate; residual quantity placed on book        |

---

## Getting Started

```bash
git clone https://github.com/peekeah/go-orderbook
cd go-orderbook
make run
```

---

## API

### Market Order
```
POST /market-order
{
  "order_type": "BUY" | "SELL",
  "quantity": 2.5
}
```

### Limit Order
```
POST /limit-order
{
  "order_type": "BUY" | "SELL",
  "quantity": 2.5,
  "rate": 118.50
}
```

### Order Book
```
GET /bids    → sorted bid ladder
GET /asks    → sorted ask ladder
```

---

## License

MIT
