package mocks

import (
	"github.com/stretchr/testify/mock"
	"time"
)

type OrderCreatedMock struct {
	mock.Mock
}

func (c *OrderCreatedMock) GetName() string {
	args := c.Called()
	return args.String(0)
}

func (c *OrderCreatedMock) GetDateTime() time.Time {
	args := c.Called()
	return args.Get(0).(time.Time)
}

func (c *OrderCreatedMock) GetPayload() interface{} {
	args := c.Called()
	return args.Get(0)
}

func (c *OrderCreatedMock) SetPayload(payload interface{}) {
	c.Called(payload)
}
