// My Simple Approach to data binding.
// for simplicity i shall start with string data items
package binding

import (
	"errors"
	"reflect"
	"sync"
)

type Boolean interface {
	DataItem
	Get() (bool, error)
	TryGet() bool
	Set(bool)
	Register(interface{})
}

type boundBoolean struct {
	str  *sync.Map
	lock sync.RWMutex
}

func (s *boundBoolean) AddListener(fn func()) {
	s.lock.Lock()
	defer s.lock.Unlock()
	listener := &listener{
		fn,
	}
	s.str.Store("Listener", listener)
}
func (s *boundBoolean) RemoveListener() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.str.Delete("Listener")
}
func (s *boundBoolean) Get() (bool, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	val, ok := s.str.Load("value")
	if ok {
		return val.(bool), nil
	} else {
		err := errors.New("no value found")
		return false, err
	}
}
func (s *boundBoolean) TryGet() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	val, ok := s.str.Load("value")
	if ok {
		return val.(bool)
	} else {
		return false
	}
}

func (s *boundBoolean) Set(str bool) {
	s.lock.Lock()
	// defer s.lock.Unlock()
	s.str.Store("value", str)
	s.lock.Unlock()
	callback, ok := s.str.Load("Listener")
	if ok {
		callback.(*listener).callback()
	}
	triggers, ok := s.str.Load("listeners")
	if ok {

		triggerMap := triggers.(map[string]interface{})
		// log.Printf("Set <%+v>", triggerMap)
		for k := range triggerMap {
			// log.Printf("[BoundBool Set] triggerMap of\n\t %s\n\t%+v\n", k, reflect.ValueOf(triggerMap[k]))
			// log.Printf("[BoundBool Set] triggerMap Bounded\n\t%+v\n", reflect.ValueOf(triggerMap[k]).MethodByName("Bounded"))
			// reflect.ValueOf(triggerMap[k]).MethodByName("Bounded").Call([]reflect.Value{})
			methodAddress := reflect.ValueOf(triggerMap[k]).MethodByName("Bounded")
			methodInterface := methodAddress.Interface().(func(interface{}))
			methodInterface(str)
		}
	}
}
func (s *boundBoolean) Register(i interface{}) {
	// we get the "listeners" from the binding.String
	// we set the Id to the string map
	// we set the Bound to the function requested
	triggers, ok := s.str.Load("listeners")
	if ok {
		triggerMap := triggers.(map[string]interface{})
		item := reflect.ValueOf(i)
		var id string
		// log.Printf("[boundString Register] item's Kind %+v\n", item.Kind())
		if item.Kind() == reflect.Ptr {
			id = item.Elem().FieldByName("Id").String()
		} else {
			id = item.FieldByName("Id").String()
		}
		// bounded := item.MethodByName("Bounded")
		// log.Printf("Registering \n\t<ID = %s>, \n\t<Bounded = %+v >\n", id, bounded)
		triggerMap[id] = i
		// try to run the function in Bound

		// if funcM.IsZero() {
		// 	log.Printf("FuncM is its Zero Value ")
		// } else {
		// 	funcM.Call([]reflect.Value{})
		// }

		// item := i.(*triggerItem)
		// triggerMap[item.Id] = item.Bound
	}
}

func NewBoolean() Boolean {
	var str sync.Map = sync.Map{}
	str.Store("value", false)
	str.Store("listeners", map[string]interface{}{})
	return &boundBoolean{
		str: &str,
	}
}
