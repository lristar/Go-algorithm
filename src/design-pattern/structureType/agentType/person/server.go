package person

type IServer interface {
	SellProduct(num int)
	GetProductNum() int
	GetProductBalance() float64
}

type Server struct {
	ProductName    string
	ProductBalance float64
	ProductNum     int
	Values         float64
}

func (s *Server) SellProduct(num int) {
	b := (s.ProductBalance) * float64(num)
	s.ProductNum -= num
	s.Values += b
}

func (s *Server) GetProductNum() int {
	return s.ProductNum
}

func (s *Server) GetProductBalance() float64 {
	return s.ProductBalance
}


