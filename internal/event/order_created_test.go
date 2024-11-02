package event

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenAnEventInput_WhenICallNewOrderCreated_ThenIShouldReceiveOrderCreatedEvent(t *testing.T) {
	event := NewOrderCreated()
	assert.Equal(t, "OrderCreated", event.GetName())
}

func TestGivenAnEventInput_WhenICallSetPayload_ThenIShouldReceivePayload(t *testing.T) {
	event := NewOrderCreated()
	event.SetPayload("test")
	assert.Equal(t, "test", event.GetPayload())
}

func TestGivenAnEventInput_WhenICallGetDateTime_ThenIShouldReceiveCurrentTime(t *testing.T) {
	event := NewOrderCreated()
	assert.NotNil(t, event.GetDateTime())
}
