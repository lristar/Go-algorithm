package proto

import "Go-algorithm/src/design-pattern/createType/protoType"

type Rectangle struct {
	id    int
	types string
}

func (r *Rectangle) GetType() string {
	return r.types
}

func (r *Rectangle) GetId() int {
	return r.id
}

func (r *Rectangle) SetId(id int) {
	r.id = id
}

func (r *Rectangle) Settypes(types string) {
	r.types = types
}

func (r *Rectangle) Clone() protoType.Shape {
	return r
}