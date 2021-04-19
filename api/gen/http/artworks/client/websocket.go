// Code generated by goa v3.3.1, DO NOT EDIT.
//
// artworks WebSocket client streaming
//
// Command:
// $ goa gen github.com/pastelnetwork/walletnode/api/design

package client

import (
	"io"

	"github.com/gorilla/websocket"
	artworks "github.com/pastelnetwork/walletnode/api/gen/artworks"
	artworksviews "github.com/pastelnetwork/walletnode/api/gen/artworks/views"
	goahttp "goa.design/goa/v3/http"
)

// ConnConfigurer holds the websocket connection configurer functions for the
// streaming endpoints in "artworks" service.
type ConnConfigurer struct {
	RegisterStatusFn goahttp.ConnConfigureFunc
}

// RegisterStatusClientStream implements the
// artworks.RegisterStatusClientStream interface.
type RegisterStatusClientStream struct {
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// NewConnConfigurer initializes the websocket connection configurer function
// with fn for all the streaming endpoints in "artworks" service.
func NewConnConfigurer(fn goahttp.ConnConfigureFunc) *ConnConfigurer {
	return &ConnConfigurer{
		RegisterStatusFn: fn,
	}
}

// Recv reads instances of "artworks.Job" from the "registerStatus" endpoint
// websocket connection.
func (s *RegisterStatusClientStream) Recv() (*artworks.Job, error) {
	var (
		rv   *artworks.Job
		body RegisterStatusResponseBody
		err  error
	)
	err = s.conn.ReadJSON(&body)
	if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
		s.conn.Close()
		return rv, io.EOF
	}
	if err != nil {
		return rv, err
	}
	res := NewRegisterStatusJobOK(&body)
	vres := &artworksviews.Job{
		Projected: res,
		View:      "default",
	}
	if err := artworksviews.ValidateJob(vres); err != nil {
		return rv, goahttp.ErrValidationError("artworks", "registerStatus", err)
	}
	return artworks.NewJob(vres), nil
}
