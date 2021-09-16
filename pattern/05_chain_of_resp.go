package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/
type Request struct {
	amount int
}

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request *Request) bool
}
type CreditCardHandler struct {
	next Handler
}

func (c *CreditCardHandler) SetNext(handler Handler) Handler {
	c.next = handler
	return handler
}
func (c *CreditCardHandler) Handle(request *Request) bool {
	if request.amount > 5000 && 10000 <= request.amount {
		fmt.Println("credit card handler processed the request")
		return true
	}
	if c.next != nil {
		return c.next.Handle(request)
	}
	fmt.Println("no handler available in credit handler")
	return false
}

type BankHandler struct {
	next Handler
}

func (b *BankHandler) SetNext(handler Handler) Handler {
	b.next = handler
	return handler
}
func (b *BankHandler) Handle(request *Request) bool {
	if request.amount < 1000 {
		fmt.Println("bank handler processed the request")
		return true
	}
	if b.next != nil {
		return b.next.Handle(request)
	}
	fmt.Println("no handler available in bank handler")
	return false
}

type LoanHandler struct {
	next Handler
}

func (l *LoanHandler) SetNext(handler Handler) Handler {
	l.next = handler
	return handler
}
func (l *LoanHandler) Handle(request *Request) bool {
	if request.amount >= 10000 {
		fmt.Println("loan handler processed the request")
		return true
	}
	if l.next != nil {
		return l.next.Handle(request)
	}
	fmt.Println("no handler available in loan handler")
	return false
}
