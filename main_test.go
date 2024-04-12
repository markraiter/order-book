package main

import (
	"testing"
)

func TestInsertOrder(t *testing.T) {
	ob := &OrderBook{}

	order := &Order{
		UserID: 1,
		Amount: 5,
		Price:  10,
		Side:   true,
	}

	ob.InsertOrder(order)

	if len(ob.BuyOrders) != 1 {
		t.Errorf("Expected BuyOrders length to be 1, got %d", len(ob.BuyOrders))
	}
}

func TestMatchOrders(t *testing.T) {
	ob := &OrderBook{}

	buyOrder := &Order{
		UserID: 1,
		Amount: 5,
		Price:  10,
		Side:   true,
	}

	sellOrder := &Order{
		UserID: 2,
		Amount: 5,
		Price:  2,
		Side:   false,
	}

	ob.InsertOrder(buyOrder)
	ob.InsertOrder(sellOrder)

	balanceChanges := ob.MatchOrders()

	if len(balanceChanges) != 4 {
		t.Errorf("Expected 4 balance changes, got %d", len(balanceChanges))
	}
}

func TestMin(t *testing.T) {
	a := int64(5)
	b := int64(10)

	result := min(a, b)

	if result != a {
		t.Errorf("Expected min of %d and %d to be %d, got %d", a, b, a, result)
	}
}
