package adaptee


type AliPay struct {

}

func NewAliPay() Adaptee{
	return &AliPay{}
}

func (a *AliPay) SpecificPay() string{
	return "使用支付宝支付"
}