// My Simple Approach to data binding.
// for simplicity i shall start with string data items
package binding

import (
	"errors"
	"fmt"
	"log"
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

type DataItem interface {
	AddListener(func())
	RemoveListener()
}
type listener struct {
	callback func()
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
	defer s.lock.Unlock()
	s.str.Store("value", str)
	callback, ok := s.str.Load("Listener")
	if ok {
		callback.(*listener).callback()
	}
	triggers, ok := s.str.Load("listeners")
	if ok {

		triggerMap := triggers.(map[string]interface{})
		log.Printf("Set <%+v>", triggerMap)
		for k := range triggerMap {
			triggerMap[k].(Boundable).Notify()

		}
	}
}
func (s *boundString) Register(i interface{}) {
	_, ok := s.str.Load("listeners")
	if ok {
		//triggerMap := triggers.(map[string]interface{})
		item := reflect.ValueOf(i).Elem()
		fmt.Printf("<Items in I % +v >\n", item)
		fmt.Printf("Item's Value is %s\n", item.FieldByName("Id").String())
		// try to run the function in Bound
		funcM := item.MethodByName("Bound")

		funcM.Call([]reflect.Value{})

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
