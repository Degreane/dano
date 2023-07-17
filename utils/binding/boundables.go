package binding

import "log"

type BoundableType interface {
	String
}

type Boundable interface {
	Notify()
}

type Bound struct {
	Widget   interface{}
	Callback func()
}

func (b *Bound) Notify() {
	b.Callback()
}
func NewBound(w interface{}, fn func()) Boundable {
	log.Printf("NewBound Registering New Binding %+v <%+v>\n ", w, fn)
	return &Bound{
		Widget:   w,
		Callback: fn,
	}
}
