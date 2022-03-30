package observer

import (
	"context"
)

var Observer = Observers{}

func init() {
	Observer = Observers{
		uniqueObservers: make(map[string]Obs),
		observers:       make(map[string][]IObserver),
	}
}

type Observers struct {
	uniqueObservers map[string]Obs
	observers       map[string][]IObserver
}

type IObserver interface {
	OnNotify(context.Context, Event) error
}

type Event struct {
	Name    string
	Message interface{}
}

// IObserver is kept as a map to avoid duplication
type Obs map[IObserver]struct{}

func (obs Observers) Publish(ctx context.Context, name string, message interface{}) error {
	// call the observers
	for _, o := range obs.observers[name] {
		err := o.OnNotify(ctx, Event{name, message})
		if err != nil {
			return err
		}
	}
	return nil
}

func (obs Observers) Register(name string, o IObserver) {
	// In case the observers
	if obs.observerRegistered(name, o) {
		return
	}

	obs.observers[name] = append(obs.observers[name], o)
	obs.storeUniqueObserver(name, o)
}

func (obs Observers) storeUniqueObserver(name string, o IObserver) {
	// call the observers
	if obs.uniqueObservers[name] == nil {
		v := map[IObserver]struct{}{}
		v[o] = struct{}{}
		obs.uniqueObservers[name] = v
	} else {
		obs.uniqueObservers[name][o] = struct{}{}
	}
}

func (obs Observers) observerRegistered(name string, o IObserver) bool {
	// call the observers
	if _, ok := obs.uniqueObservers[name]; !ok {
		return false
	}

	if _, ok := obs.uniqueObservers[name][o]; !ok {
		return false
	}

	return true
}
