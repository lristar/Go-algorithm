package person

import (
	"errors"
)

type IClient interface {
}

type Client struct {
	product map[string]int
	values  float64
}

func NewClient(value float64) *Client {
	return &Client{
		product: make(map[string]int),
		values:  value,
	}
}

func (c *Client) BuyProduct(product string, num int, productBalance float64) error {
	all := productBalance * float64(num)
	if c.values > all {
		v, ok := c.product[product]
		if !ok {
			c.product[product] = num
		}
		oldNum := v
		c.product[product] = oldNum + num
		c.values -= all
		return nil
	}
	return errors.New("不够钱")
}
