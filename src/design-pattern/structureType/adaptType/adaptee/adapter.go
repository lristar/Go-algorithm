package adaptee

// 适配的集合
var AllPay = map[string]Adaptee{
	"wechat":NewWechatPay(),
	"alipay":NewAliPay(),
	"aggregationpay":NewAggregationPay(),
}


// 目标接口
type Target interface {
	RequestPay(types string) string
}

// 被适配的目标接口
type Adaptee interface {
	SpecificPay() string
}

type Adapter struct {
	adatetors map[string]Adaptee
}

func NewAdapter()Target{
	return &Adapter{AllPay}
}

func (a *Adapter) RequestPay(types string) string{
	if target , ok :=a.adatetors[types];ok{
		return target.SpecificPay()
	}
	return "没有适配的东西"
}


