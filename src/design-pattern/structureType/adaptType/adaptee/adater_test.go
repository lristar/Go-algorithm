package adaptee

import (
	"fmt"
	"testing"
)

func TestAdapter_Request(t *testing.T) {
	target := NewAdapter()
	fmt.Println(target.RequestPay("aggregationpay"))
}