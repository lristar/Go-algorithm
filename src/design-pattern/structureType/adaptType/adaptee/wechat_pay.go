package adaptee


type WechatPay struct {

}

func NewWechatPay() Adaptee{
	return &WechatPay{}
}

func (a *WechatPay) SpecificPay() string{
	return "使用微信支付"
}