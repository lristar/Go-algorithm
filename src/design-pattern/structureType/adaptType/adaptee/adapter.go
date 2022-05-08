package adaptee

// 目标接口
type Target interface {
	RequestPay(types string) string
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


