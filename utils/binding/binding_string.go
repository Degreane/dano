// My Simple Approach to data binding.
// for simplicity i shall start with string data items
package binding

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type String interface {
	DataItem
	Get() (string, error)
	TryGet() string
	Set(string)
	Register(interface{})
}
type boundString struct {
	str  *sync.Map
	lock sync.RWMutex
}

func (s *boundString) AddListener(fn func()) {
	s.lock.Lock()
	defer s.lock.Unlock()
	listener := &listener{
		fn,
	}
	s.str.Store("Listener", listener)
}
func (s *boundString) RemoveListener() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.str.Delete("Listener")
}
func (s *boundString) Get() (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	val, ok := s.str.Load("value")
	if ok {
		return fmt.Sprintf("%v", val), nil
	} else {
		err := errors.New("no value found")
		return "", err
	}
}
func (s *boundString) TryGet() string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	val, ok := s.str.Load("value")
	if ok {
		return fmt.Sprintf("%v", val)
	} else {
		return ""
	}
}

func (s *boundString) Set(str string) {
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
			// log.Printf("[BoundString Set] triggerMap of\n\t %s\n\t%+v\n", k, reflect.ValueOf(triggerMap[k]))
			// log.Printf("[BoundString Set] triggerMap Bounded\n\t%+v\n", reflect.ValueOf(triggerMap[k]).MethodByName("Bounded"))
			// ins:=make([]reflect.Value,1)
			// ins[0]=interface{}{s.TryGet()}
			methodAddress := reflect.ValueOf(triggerMap[k]).MethodByName("Bounded")
			methodInterface := methodAddress.Interface().(func(interface{}))
			methodInterface(str)

		}
	}
}
func (s *boundString) Register(i interface{}) {
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

func NewString() String {
	var str sync.Map = sync.Map{}
	str.Store("value", "")
	str.Store("listeners", map[string]interface{}{})
	return &boundString{
		str: &str,
	}
}
