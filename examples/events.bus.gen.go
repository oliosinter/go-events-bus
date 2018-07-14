// Code generated by events-bus-gen. DO NOT EDIT.
// source: events.go

package main

import (
	productModel "github.com/oliosinter/go-events-bus/examples/models"
	"reflect"
	"sync"
)

type (
	// EventsUserRegistrationHandlerFunc is a listener handler function for event 'UserRegistration'
	EventsUserRegistrationHandlerFunc func(arg0 *UserInfo)
	// EventsProductUpdateHandlerFunc is a listener handler function for event 'ProductUpdate'
	EventsProductUpdateHandlerFunc func(arg0 *productModel.Product)
	// EventsCombinedEventHandlerFunc is a listener handler function for event 'CombinedEvent'
	EventsCombinedEventHandlerFunc func(arg0 UserInfo, arg1 productModel.Product)
)

// NewEventsEmitter creates new EventsEmitter
func NewEventsEmitter() *EventsEmitter {
	return &EventsEmitter{}
}

// EventsEmitter implements events listener and events emitter operations
// for events UserRegistration, ProductUpdate, CombinedEvent
type EventsEmitter struct {
	lockUserRegistration sync.RWMutex
	onUserRegistration   []EventsUserRegistrationHandlerFunc
	lockProductUpdate sync.RWMutex
	onProductUpdate   []EventsProductUpdateHandlerFunc
	lockCombinedEvent sync.RWMutex
	onCombinedEvent   []EventsCombinedEventHandlerFunc
}

// EventsBus is a client side of event bus that allows subscribe to
// UserRegistration, ProductUpdate, CombinedEvent events
type EventsBus interface {
	Events
	// UserRegistration adds event listener for event 'UserRegistration'
	OnUserRegistration(handler EventsUserRegistrationHandlerFunc)
	// RemoveUserRegistration excludes event listener
	RemoveUserRegistration(handler EventsUserRegistrationHandlerFunc)
	// ProductUpdate adds event listener for event 'ProductUpdate'
	OnProductUpdate(handler EventsProductUpdateHandlerFunc)
	// RemoveProductUpdate excludes event listener
	RemoveProductUpdate(handler EventsProductUpdateHandlerFunc)
	// CombinedEvent adds event listener for event 'CombinedEvent'
	OnCombinedEvent(handler EventsCombinedEventHandlerFunc)
	// RemoveCombinedEvent excludes event listener
	RemoveCombinedEvent(handler EventsCombinedEventHandlerFunc)
}

// OnUserRegistration adds event listener for event 'UserRegistration'
func (bus *EventsEmitter) OnUserRegistration(handler EventsUserRegistrationHandlerFunc) {
	bus.lockUserRegistration.Lock()
	defer bus.lockUserRegistration.Unlock()
	bus.onUserRegistration = append(bus.onUserRegistration, handler)
}

// RemoveUserRegistration excludes event listener
func (bus *EventsEmitter) RemoveUserRegistration(handler EventsUserRegistrationHandlerFunc) {
	bus.lockUserRegistration.Lock()
	defer bus.lockUserRegistration.Unlock()
	var res []EventsUserRegistrationHandlerFunc
	refVal := reflect.ValueOf(handler).Pointer()
	for _, f := range bus.onUserRegistration {
		if reflect.ValueOf(f).Pointer() != refVal {
			res = append(res, f)
		}
	}
	bus.onUserRegistration = res
}

// UserRegistration emits event with same name
func (bus *EventsEmitter) UserRegistration(arg0 *UserInfo) {
	bus.lockUserRegistration.RLock()
	defer bus.lockUserRegistration.RUnlock()
	for _, f := range bus.onUserRegistration {
		f(arg0)
	}
}

// OnProductUpdate adds event listener for event 'ProductUpdate'
func (bus *EventsEmitter) OnProductUpdate(handler EventsProductUpdateHandlerFunc) {
	bus.lockProductUpdate.Lock()
	defer bus.lockProductUpdate.Unlock()
	bus.onProductUpdate = append(bus.onProductUpdate, handler)
}

// RemoveProductUpdate excludes event listener
func (bus *EventsEmitter) RemoveProductUpdate(handler EventsProductUpdateHandlerFunc) {
	bus.lockProductUpdate.Lock()
	defer bus.lockProductUpdate.Unlock()
	var res []EventsProductUpdateHandlerFunc
	refVal := reflect.ValueOf(handler).Pointer()
	for _, f := range bus.onProductUpdate {
		if reflect.ValueOf(f).Pointer() != refVal {
			res = append(res, f)
		}
	}
	bus.onProductUpdate = res
}

// ProductUpdate emits event with same name
func (bus *EventsEmitter) ProductUpdate(arg0 *productModel.Product) {
	bus.lockProductUpdate.RLock()
	defer bus.lockProductUpdate.RUnlock()
	for _, f := range bus.onProductUpdate {
		f(arg0)
	}
}

// OnCombinedEvent adds event listener for event 'CombinedEvent'
func (bus *EventsEmitter) OnCombinedEvent(handler EventsCombinedEventHandlerFunc) {
	bus.lockCombinedEvent.Lock()
	defer bus.lockCombinedEvent.Unlock()
	bus.onCombinedEvent = append(bus.onCombinedEvent, handler)
}

// RemoveCombinedEvent excludes event listener
func (bus *EventsEmitter) RemoveCombinedEvent(handler EventsCombinedEventHandlerFunc) {
	bus.lockCombinedEvent.Lock()
	defer bus.lockCombinedEvent.Unlock()
	var res []EventsCombinedEventHandlerFunc
	refVal := reflect.ValueOf(handler).Pointer()
	for _, f := range bus.onCombinedEvent {
		if reflect.ValueOf(f).Pointer() != refVal {
			res = append(res, f)
		}
	}
	bus.onCombinedEvent = res
}

// CombinedEvent emits event with same name
func (bus *EventsEmitter) CombinedEvent(arg0 UserInfo, arg1 productModel.Product) {
	bus.lockCombinedEvent.RLock()
	defer bus.lockCombinedEvent.RUnlock()
	for _, f := range bus.onCombinedEvent {
		f(arg0, arg1)
	}
}
