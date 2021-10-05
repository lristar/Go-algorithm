package item

type Item interface {
	Packing()Packing
	Getname()string
	Getprice()float64

}

type Packing interface {
	Pack()string
}