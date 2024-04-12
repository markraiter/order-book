package main

import (
	"fmt"
	"sort"
)

type Order struct {
	UserID int64
	Amount int64
	Price  int64
	Side   bool // true for buy, false for sell
}

type BalanceChange struct {
	UserID   int64
	Value    int64
	Currency string
}

type OrderBook struct {
	BuyOrders  []*Order
	SellOrders []*Order
}

func (ob *OrderBook) InsertOrder(order *Order) {
	if order.Side {
		ob.BuyOrders = append(ob.BuyOrders, order)
		sort.Slice(ob.BuyOrders, func(i, j int) bool {
			return ob.BuyOrders[i].Price > ob.BuyOrders[j].Price
		})
	} else {
		ob.SellOrders = append(ob.SellOrders, order)
		sort.Slice(ob.SellOrders, func(i, j int) bool {
			return ob.SellOrders[i].Price < ob.SellOrders[j].Price
		})
	}
}

func (ob *OrderBook) MatchOrders() []*BalanceChange {
	var balanceChanges []*BalanceChange

	for len(ob.BuyOrders) > 0 && len(ob.SellOrders) > 0 {
		buyOrder := ob.BuyOrders[0]
		sellOrder := ob.SellOrders[0]

		if buyOrder.Price >= sellOrder.Price {
			amount := min(buyOrder.Amount, sellOrder.Amount)
			buyOrder.Amount -= amount
			sellOrder.Amount -= amount

			balanceChanges = append(balanceChanges, &BalanceChange{
				UserID:   buyOrder.UserID,
				Value:    -amount * buyOrder.Price,
				Currency: "USD",
			}, &BalanceChange{
				UserID:   sellOrder.UserID,
				Value:    amount * sellOrder.Price,
				Currency: "USD",
			}, &BalanceChange{
				UserID:   buyOrder.UserID,
				Value:    amount,
				Currency: "UAH",
			}, &BalanceChange{
				UserID:   sellOrder.UserID,
				Value:    -amount,
				Currency: "UAH",
			})

			if buyOrder.Amount == 0 {
				ob.BuyOrders = ob.BuyOrders[1:]
			}
			if sellOrder.Amount == 0 {
				ob.SellOrders = ob.SellOrders[1:]
			}
		} else {
			break
		}
	}

	return balanceChanges
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	ob := &OrderBook{}

	fmt.Println("Enter orders (UserID, Amount, Price, Side). Enter -1 for UserID to stop.")

	for {
		var userID, amount, price int64
		var side bool

		fmt.Print("UserID: ")
		fmt.Scan(&userID)
		if userID == -1 {
			break
		}

		fmt.Print("Amount: ")
		fmt.Scan(&amount)

		fmt.Print("Price: ")
		fmt.Scan(&price)

		fmt.Print("Side (1 for buy, 0 for sell): ")
		fmt.Scan(&side)

		order := &Order{
			UserID: userID,
			Amount: amount,
			Price:  price,
			Side:   side,
		}

		ob.InsertOrder(order)
	}

	for _, bc := range ob.MatchOrders() {
		fmt.Printf("UserID: %d, Value: %d, Currency: %s\n", bc.UserID, bc.Value, bc.Currency)
	}
}
