package main

import "fmt"

type PaymentContext struct {
	Name, CardID string
	Money        int
	payment      PaymentStrategy
}

func (p *PaymentContext) Pay() {
	p.payment.Pay(p)
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

type Cash struct {
}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Println("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

type Bank struct {
}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Println("Pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)

}

func main() {
	cash := PaymentContext{Name: "Ada", CardID: "", Money: 123, payment: &Cash{}}
	cash.Pay()
	bank := PaymentContext{Name: "bank", CardID: "", Money: 456, payment: &Bank{}}
	bank.Pay()
}
