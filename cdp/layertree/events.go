package layertree

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	cdp "github.com/rjeczalik/chromedp/cdp"
	"github.com/rjeczalik/chromedp/cdp/dom"
)

// EventLayerTreeDidChange [no description].
type EventLayerTreeDidChange struct {
	Layers []*Layer `json:"layers,omitempty"` // Layer tree, absent if not in the comspositing mode.
}

// EventLayerPainted [no description].
type EventLayerPainted struct {
	LayerID LayerID   `json:"layerId"` // The id of the painted layer.
	Clip    *dom.Rect `json:"clip"`    // Clip rectangle.
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventLayerTreeLayerTreeDidChange,
	cdp.EventLayerTreeLayerPainted,
}
