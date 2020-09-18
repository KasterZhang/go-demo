package strategy

import (
	"fmt"
)

//design model: startegy model

//POS POS machine
type POS interface {
	pay(int64)
}

type seller struct {
	pos POS
}

type alipay struct{}

type wechat struct{}

type stripe struct{}

func (p *alipay) pay(amount int64) {
	fmt.Println("支付宝到账:", amount/100, "元")
}

func (p *wechat) pay(amount int64) {
	fmt.Println("微信到账:", amount/100, "元")
}

func (p *stripe) pay(amount int64) {
	fmt.Println("stripe到账:", amount/100, "元")
}

func (p *seller) changePaymentype(pos POS) {
	p.pos = pos
}

func (p *seller) swipe(amount int64) {
	p.pos.pay(amount)
}
