package main

import "fmt"

type PayMethod02 interface {
	Pay(int)
}

type CreditCard02 struct {
	balance int
	limit   int
}

func (c *CreditCard02) Pay(amout int) {
	if c.balance < amout {
		fmt.Println("余额不足")
		return
	}
	c.balance -= amout
}

// 空接口在接收参数方面，有很大的灵活性
func anyParam(param interface{}) {
	fmt.Println("param: ", param)
}

func main02() {
	c := CreditCard02{balance: 100, limit: 1000}
	c.Pay(200)
	var a PayMethod02 = &c
	fmt.Println("a.Pay: ", a)

	var b interface{} = &c
	fmt.Println("b: ", b)

	anyParam(c)
	anyParam(1)
	anyParam("123")
	anyParam(a)
}
