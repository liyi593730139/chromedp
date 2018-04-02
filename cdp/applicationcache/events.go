package applicationcache

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	cdp "github.com/rjeczalik/chromedp/cdp"
)

// EventApplicationCacheStatusUpdated [no description].
type EventApplicationCacheStatusUpdated struct {
	FrameID     cdp.FrameID `json:"frameId"`     // Identifier of the frame containing document whose application cache updated status.
	ManifestURL string      `json:"manifestURL"` // Manifest URL.
	Status      int64       `json:"status"`      // Updated application cache status.
}

// EventNetworkStateUpdated [no description].
type EventNetworkStateUpdated struct {
	IsNowOnline bool `json:"isNowOnline"`
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventApplicationCacheApplicationCacheStatusUpdated,
	cdp.EventApplicationCacheNetworkStateUpdated,
}
