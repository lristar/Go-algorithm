package adaptee


type Aggregation struct {

}

func NewAggregationPay() Adaptee{
	return &Aggregation{}
}

func (a *Aggregation) SpecificPay() string{
	return "使用聚合支付支付"
}