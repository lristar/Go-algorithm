package adaptee


// 被适配的目标接口
type Adaptee interface {
	SpecificPay() string
}