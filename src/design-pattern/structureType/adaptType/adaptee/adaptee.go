package adaptee

// 适配的集合
var AllPay = map[string]Adaptee{
	"wechat":NewWechatPay(),
	"alipay":NewAliPay(),
	"aggregationpay":NewAggregationPay(),
}
// 被适配的目标接口
type Adaptee interface {
	SpecificPay() string
}