package tracing

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	cdp "github.com/rjeczalik/chromedp/cdp"
	"github.com/rjeczalik/chromedp/cdp/io"
	"github.com/mailru/easyjson"
)

// EventDataCollected contains an bucket of collected trace events. When
// tracing is stopped collected events will be send as a sequence of
// dataCollected events followed by tracingComplete event.
type EventDataCollected struct {
	Value []easyjson.RawMessage `json:"value"`
}

// EventTracingComplete signals that tracing is stopped and there is no trace
// buffers pending flush, all data were delivered via dataCollected events.
type EventTracingComplete struct {
	Stream io.StreamHandle `json:"stream,omitempty"` // A handle of the stream that holds resulting trace data.
}

// EventBufferUsage [no description].
type EventBufferUsage struct {
	PercentFull float64 `json:"percentFull,omitempty"` // A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
	EventCount  float64 `json:"eventCount,omitempty"`  // An approximate number of events in the trace log.
	Value       float64 `json:"value,omitempty"`       // A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventTracingDataCollected,
	cdp.EventTracingTracingComplete,
	cdp.EventTracingBufferUsage,
}
