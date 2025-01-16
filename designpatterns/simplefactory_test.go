package designpatterns

import (
	"fmt"
	"testing"
)

type PayType int

const (
	AlipayType PayType   = 1
	WechatpayType PayType = 2
)

type Pay interface {
	PayPage(price int) string
}

type Alipay struct {
}

func (Alipay) PayPage(price int) string {
	return fmt.Sprintf("alipay: %d", price)
}

type WechatPay struct {
}

func (WechatPay) PayPage(price int) string {
	return fmt.Sprintf("Wechatpay: %d", price)
}

func NewPayPage(typ PayType) Pay {
	switch typ {
	case AlipayType:
		return Alipay{}
	case WechatpayType:
		return WechatPay{}
	default:
		return nil
	}
}

func TestSimpleFactory(t *testing.T) {
	pay := NewPayPage(AlipayType)
	t.Error(pay.PayPage(12))
	pay = NewPayPage(WechatpayType)
	t.Error(pay.PayPage(89))
}
