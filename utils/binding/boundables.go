package binding

import "log"

type DataItem interface {
	AddListener(func())
	RemoveListener()
}
type listener struct {
	callback func()
}
type BoundableStringType interface {
	String
}
type BoundableBoolType interface {
	Boolean
}

type Boundable interface {
	Notify()
}

type Bound struct {
	Widget   interface{}
	Callback func()
}

func (b *Bound) Notify() {
	log.Printf("Calling Notify")
	b.Callback()
}
func NewBound(w interface{}, fn func()) Boundable {
	// log.Printf("[Boundable NewBound] NewBound Boundable Registering New Binding %+v <%+v>\n ", w, fn)
	return &Bound{
		Widget:   w,
		Callback: fn,
	}
}
