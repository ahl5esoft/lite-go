package mongosvc

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"go.mongodb.org/mongo-driver/event"
)

type monitor struct {
	span opentracing.Span
}

func (m *monitor) failed(_ context.Context, evt *event.CommandFailedEvent) {
	if m.isExec(evt.CommandName) || evt == nil {
		return
	}

	defer m.span.Finish()

	m.span.LogFields(
		log.String("failed", evt.Failure),
	)
}

func (m monitor) isExec(command string) bool {
	return command == "ping"
}

func (m *monitor) started(ctx context.Context, evt *event.CommandStartedEvent) {
	if m.isExec(evt.CommandName) || evt == nil {
		return
	}

	var parentCtx opentracing.SpanContext
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan != nil {
		parentCtx = parentSpan.Context()
	}

	m.span = opentracing.GlobalTracer().StartSpan(
		"mongo",
		opentracing.ChildOf(parentCtx),
	)
	m.span.SetTag("started", map[string]string{
		"conn":    evt.ConnectionID,
		"cmd":     evt.Command.String(),
		"cmdName": evt.CommandName,
	})
}

func (m *monitor) succeeded(_ context.Context, evt *event.CommandSucceededEvent) {
	if m.isExec(evt.CommandName) || evt == nil {
		return
	}

	defer m.span.Finish()

	m.span.LogFields(
		log.String(
			"succeeded",
			evt.Reply.String(),
		),
	)
}
