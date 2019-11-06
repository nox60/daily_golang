package main

type Customer interface {
	Accept(Visitor)
}

type Visitor interface {
	Visit(Customer)
}
