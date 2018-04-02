// Package audits provides the Chrome Debugging Protocol
// commands, types, and events for the Audits domain.
//
// Audits domain allows investigation of page violations and possible
// improvements.
//
// Generated by the chromedp-gen command.
package audits

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"context"

	cdp "github.com/rjeczalik/chromedp/cdp"
	"github.com/rjeczalik/chromedp/cdp/network"
)

// GetEncodedResponseParams returns the response body and size if it were
// re-encoded with the specified settings. Only applies to images.
type GetEncodedResponseParams struct {
	RequestID network.RequestID          `json:"requestId"`          // Identifier of the network request to get content for.
	Encoding  GetEncodedResponseEncoding `json:"encoding"`           // The encoding to use.
	Quality   float64                    `json:"quality,omitempty"`  // The quality of the encoding (0-1). (defaults to 1)
	SizeOnly  bool                       `json:"sizeOnly,omitempty"` // Whether to only return the size information (defaults to false).
}

// GetEncodedResponse returns the response body and size if it were
// re-encoded with the specified settings. Only applies to images.
//
// parameters:
//   requestID - Identifier of the network request to get content for.
//   encoding - The encoding to use.
func GetEncodedResponse(requestID network.RequestID, encoding GetEncodedResponseEncoding) *GetEncodedResponseParams {
	return &GetEncodedResponseParams{
		RequestID: requestID,
		Encoding:  encoding,
	}
}

// WithQuality the quality of the encoding (0-1). (defaults to 1).
func (p GetEncodedResponseParams) WithQuality(quality float64) *GetEncodedResponseParams {
	p.Quality = quality
	return &p
}

// WithSizeOnly whether to only return the size information (defaults to
// false).
func (p GetEncodedResponseParams) WithSizeOnly(sizeOnly bool) *GetEncodedResponseParams {
	p.SizeOnly = sizeOnly
	return &p
}

// GetEncodedResponseReturns return values.
type GetEncodedResponseReturns struct {
	Body         string `json:"body,omitempty"`         // The encoded body as a base64 string. Omitted if sizeOnly is true.
	OriginalSize int64  `json:"originalSize,omitempty"` // Size before re-encoding.
	EncodedSize  int64  `json:"encodedSize,omitempty"`  // Size after re-encoding.
}

// Do executes Audits.getEncodedResponse against the provided context and
// target handler.
//
// returns:
//   body - The encoded body as a base64 string. Omitted if sizeOnly is true.
//   originalSize - Size before re-encoding.
//   encodedSize - Size after re-encoding.
func (p *GetEncodedResponseParams) Do(ctxt context.Context, h cdp.Handler) (body string, originalSize int64, encodedSize int64, err error) {
	// execute
	var res GetEncodedResponseReturns
	err = h.Execute(ctxt, cdp.CommandAuditsGetEncodedResponse, p, &res)
	if err != nil {
		return "", 0, 0, err
	}

	return res.Body, res.OriginalSize, res.EncodedSize, nil
}
