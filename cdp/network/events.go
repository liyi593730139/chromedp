package network

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	cdp "github.com/rjeczalik/chromedp/cdp"
	"github.com/rjeczalik/chromedp/cdp/page"
)

// EventResourceChangedPriority fired when resource loading priority is
// changed.
type EventResourceChangedPriority struct {
	RequestID   RequestID          `json:"requestId"`   // Request identifier.
	NewPriority ResourcePriority   `json:"newPriority"` // New priority
	Timestamp   *cdp.MonotonicTime `json:"timestamp"`   // Timestamp.
}

// EventRequestWillBeSent fired when page is about to send HTTP request.
type EventRequestWillBeSent struct {
	RequestID        RequestID           `json:"requestId"`                  // Request identifier.
	LoaderID         cdp.LoaderID        `json:"loaderId"`                   // Loader identifier. Empty string if the request is fetched from worker.
	DocumentURL      string              `json:"documentURL"`                // URL of the document this request is loaded for.
	Request          *Request            `json:"request"`                    // Request data.
	Timestamp        *cdp.MonotonicTime  `json:"timestamp"`                  // Timestamp.
	WallTime         *cdp.TimeSinceEpoch `json:"wallTime"`                   // Timestamp.
	Initiator        *Initiator          `json:"initiator"`                  // Request initiator.
	RedirectResponse *Response           `json:"redirectResponse,omitempty"` // Redirect response data.
	Type             page.ResourceType   `json:"type,omitempty"`             // Type of this resource.
	FrameID          cdp.FrameID         `json:"frameId,omitempty"`          // Frame identifier.
}

// EventRequestServedFromCache fired if request ended up loading from cache.
type EventRequestServedFromCache struct {
	RequestID RequestID `json:"requestId"` // Request identifier.
}

// EventResponseReceived fired when HTTP response is available.
type EventResponseReceived struct {
	RequestID RequestID          `json:"requestId"`         // Request identifier.
	LoaderID  cdp.LoaderID       `json:"loaderId"`          // Loader identifier. Empty string if the request is fetched from worker.
	Timestamp *cdp.MonotonicTime `json:"timestamp"`         // Timestamp.
	Type      page.ResourceType  `json:"type"`              // Resource type.
	Response  *Response          `json:"response"`          // Response data.
	FrameID   cdp.FrameID        `json:"frameId,omitempty"` // Frame identifier.
}

// EventDataReceived fired when data chunk was received over the network.
type EventDataReceived struct {
	RequestID         RequestID          `json:"requestId"`         // Request identifier.
	Timestamp         *cdp.MonotonicTime `json:"timestamp"`         // Timestamp.
	DataLength        int64              `json:"dataLength"`        // Data chunk length.
	EncodedDataLength int64              `json:"encodedDataLength"` // Actual bytes received (might be less than dataLength for compressed encodings).
}

// EventLoadingFinished fired when HTTP request has finished loading.
type EventLoadingFinished struct {
	RequestID         RequestID          `json:"requestId"`         // Request identifier.
	Timestamp         *cdp.MonotonicTime `json:"timestamp"`         // Timestamp.
	EncodedDataLength float64            `json:"encodedDataLength"` // Total number of bytes received for this request.
}

// EventLoadingFailed fired when HTTP request has failed to load.
type EventLoadingFailed struct {
	RequestID     RequestID          `json:"requestId"`               // Request identifier.
	Timestamp     *cdp.MonotonicTime `json:"timestamp"`               // Timestamp.
	Type          page.ResourceType  `json:"type"`                    // Resource type.
	ErrorText     string             `json:"errorText"`               // User friendly error message.
	Canceled      bool               `json:"canceled,omitempty"`      // True if loading was canceled.
	BlockedReason BlockedReason      `json:"blockedReason,omitempty"` // The reason why loading was blocked, if any.
}

// EventWebSocketWillSendHandshakeRequest fired when WebSocket is about to
// initiate handshake.
type EventWebSocketWillSendHandshakeRequest struct {
	RequestID RequestID           `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime  `json:"timestamp"` // Timestamp.
	WallTime  *cdp.TimeSinceEpoch `json:"wallTime"`  // UTC Timestamp.
	Request   *WebSocketRequest   `json:"request"`   // WebSocket request data.
}

// EventWebSocketHandshakeResponseReceived fired when WebSocket handshake
// response becomes available.
type EventWebSocketHandshakeResponseReceived struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
	Response  *WebSocketResponse `json:"response"`  // WebSocket response data.
}

// EventWebSocketCreated fired upon WebSocket creation.
type EventWebSocketCreated struct {
	RequestID RequestID  `json:"requestId"`           // Request identifier.
	URL       string     `json:"url"`                 // WebSocket request URL.
	Initiator *Initiator `json:"initiator,omitempty"` // Request initiator.
}

// EventWebSocketClosed fired when WebSocket is closed.
type EventWebSocketClosed struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
}

// EventWebSocketFrameReceived fired when WebSocket frame is received.
type EventWebSocketFrameReceived struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
	Response  *WebSocketFrame    `json:"response"`  // WebSocket response data.
}

// EventWebSocketFrameError fired when WebSocket frame error occurs.
type EventWebSocketFrameError struct {
	RequestID    RequestID          `json:"requestId"`    // Request identifier.
	Timestamp    *cdp.MonotonicTime `json:"timestamp"`    // Timestamp.
	ErrorMessage string             `json:"errorMessage"` // WebSocket frame error message.
}

// EventWebSocketFrameSent fired when WebSocket frame is sent.
type EventWebSocketFrameSent struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
	Response  *WebSocketFrame    `json:"response"`  // WebSocket response data.
}

// EventEventSourceMessageReceived fired when EventSource message is
// received.
type EventEventSourceMessageReceived struct {
	RequestID RequestID          `json:"requestId"` // Request identifier.
	Timestamp *cdp.MonotonicTime `json:"timestamp"` // Timestamp.
	EventName string             `json:"eventName"` // Message type.
	EventID   string             `json:"eventId"`   // Message identifier.
	Data      string             `json:"data"`      // Message content.
}

// EventRequestIntercepted details of an intercepted HTTP request, which must
// be either allowed, blocked, modified or mocked.
type EventRequestIntercepted struct {
	InterceptionID      InterceptionID    `json:"interceptionId"` // Each request the page makes will have a unique id, however if any redirects are encountered while processing that fetch, they will be reported with the same id as the original fetch. Likewise if HTTP authentication is needed then the same fetch id will be used.
	Request             *Request          `json:"request"`
	FrameID             cdp.FrameID       `json:"frameId"`                       // The id of the frame that initiated the request.
	ResourceType        page.ResourceType `json:"resourceType"`                  // How the requested resource will be used.
	IsNavigationRequest bool              `json:"isNavigationRequest"`           // Whether this is a navigation request, which can abort the navigation completely.
	RedirectURL         string            `json:"redirectUrl,omitempty"`         // Redirect location, only sent if a redirect was intercepted.
	AuthChallenge       *AuthChallenge    `json:"authChallenge,omitempty"`       // Details of the Authorization Challenge encountered. If this is set then continueInterceptedRequest must contain an authChallengeResponse.
	ResponseErrorReason ErrorReason       `json:"responseErrorReason,omitempty"` // Response error if intercepted at response stage or if redirect occurred while intercepting request.
	ResponseStatusCode  int64             `json:"responseStatusCode,omitempty"`  // Response code if intercepted at response stage or if redirect occurred while intercepting request or auth retry occurred.
	ResponseHeaders     Headers           `json:"responseHeaders,omitempty"`     // Response headers if intercepted at the response stage or if redirect occurred while intercepting request or auth retry occurred.
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventNetworkResourceChangedPriority,
	cdp.EventNetworkRequestWillBeSent,
	cdp.EventNetworkRequestServedFromCache,
	cdp.EventNetworkResponseReceived,
	cdp.EventNetworkDataReceived,
	cdp.EventNetworkLoadingFinished,
	cdp.EventNetworkLoadingFailed,
	cdp.EventNetworkWebSocketWillSendHandshakeRequest,
	cdp.EventNetworkWebSocketHandshakeResponseReceived,
	cdp.EventNetworkWebSocketCreated,
	cdp.EventNetworkWebSocketClosed,
	cdp.EventNetworkWebSocketFrameReceived,
	cdp.EventNetworkWebSocketFrameError,
	cdp.EventNetworkWebSocketFrameSent,
	cdp.EventNetworkEventSourceMessageReceived,
	cdp.EventNetworkRequestIntercepted,
}
