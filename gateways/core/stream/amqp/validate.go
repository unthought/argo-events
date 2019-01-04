package amqp

import (
	"context"
	"fmt"
	"github.com/argoproj/argo-events/gateways"
)

// ValidateEventSource validates gateway event source
func (ese *AMQPEventSourceExecutor) ValidateEventSource(ctx context.Context, es *gateways.EventSource) (*gateways.ValidEventSource, error) {
	v := &gateways.ValidEventSource{}
	a, err := parseEventSource(es.Data)
	if err != nil {
		return v, gateways.ErrEventSourceParseFailed
	}
	if a == nil {
		return v, fmt.Errorf("%+v, configuration must be non empty", gateways.ErrInvalidEventSource)
	}
	if a.URL == "" {
		return v, fmt.Errorf("%+v, url must be specified", gateways.ErrInvalidEventSource)
	}
	if a.RoutingKey == "" {
		return v, fmt.Errorf("%+v, routing key must be specified", gateways.ErrInvalidEventSource)
	}
	if a.ExchangeName == "" {
		return v, fmt.Errorf("%+v, exchange name must be specified", gateways.ErrInvalidEventSource)
	}
	if a.ExchangeType == "" {
		return v, fmt.Errorf("%+v, exchange type must be specified", gateways.ErrInvalidEventSource)
	}
	return v, nil
}
