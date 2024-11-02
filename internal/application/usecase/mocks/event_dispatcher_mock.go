package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/tiagoncardoso/fc/pge/clean-arch/pkg/events"
)

type EventDispatcherMock struct {
	mock.Mock
}

func (e *EventDispatcherMock) Dispatch(event events.EventInterface) error {
	args := e.Called(event)
	return args.Error(0)
}

func (e *EventDispatcherMock) Register(eventName string, handler events.EventHandlerInterface) error {
	args := e.Called(eventName, handler)
	return args.Error(0)
}

func (e *EventDispatcherMock) Remove(eventName string, handler events.EventHandlerInterface) error {
	args := e.Called(eventName, handler)
	return args.Error(0)
}

func (e *EventDispatcherMock) Has(eventName string, handler events.EventHandlerInterface) bool {
	args := e.Called(eventName, handler)
	return args.Bool(0)
}

func (e *EventDispatcherMock) Clear() {
	e.Called()
}
