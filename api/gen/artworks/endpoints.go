// Code generated by goa v3.3.1, DO NOT EDIT.
//
// artworks endpoints
//
// Command:
// $ goa gen github.com/pastelnetwork/walletnode/api/design

package artworks

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "artworks" service endpoints.
type Endpoints struct {
	Register         goa.Endpoint
	RegisterJobState goa.Endpoint
	RegisterJob      goa.Endpoint
	RegisterJobs     goa.Endpoint
	UploadImage      goa.Endpoint
}

// RegisterJobStateEndpointInput holds both the payload and the server stream
// of the "registerJobState" method.
type RegisterJobStateEndpointInput struct {
	// Payload is the method payload.
	Payload *RegisterJobStatePayload
	// Stream is the server stream used by the "registerJobState" method to send
	// data.
	Stream RegisterJobStateServerStream
}

// NewEndpoints wraps the methods of the "artworks" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Register:         NewRegisterEndpoint(s),
		RegisterJobState: NewRegisterJobStateEndpoint(s),
		RegisterJob:      NewRegisterJobEndpoint(s),
		RegisterJobs:     NewRegisterJobsEndpoint(s),
		UploadImage:      NewUploadImageEndpoint(s),
	}
}

// Use applies the given middleware to all the "artworks" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Register = m(e.Register)
	e.RegisterJobState = m(e.RegisterJobState)
	e.RegisterJob = m(e.RegisterJob)
	e.RegisterJobs = m(e.RegisterJobs)
	e.UploadImage = m(e.UploadImage)
}

// NewRegisterEndpoint returns an endpoint function that calls the method
// "register" of service "artworks".
func NewRegisterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RegisterPayload)
		res, err := s.Register(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedRegisterResult(res, "default")
		return vres, nil
	}
}

// NewRegisterJobStateEndpoint returns an endpoint function that calls the
// method "registerJobState" of service "artworks".
func NewRegisterJobStateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		ep := req.(*RegisterJobStateEndpointInput)
		return nil, s.RegisterJobState(ctx, ep.Payload, ep.Stream)
	}
}

// NewRegisterJobEndpoint returns an endpoint function that calls the method
// "registerJob" of service "artworks".
func NewRegisterJobEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RegisterJobPayload)
		res, err := s.RegisterJob(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedJob(res, "default")
		return vres, nil
	}
}

// NewRegisterJobsEndpoint returns an endpoint function that calls the method
// "registerJobs" of service "artworks".
func NewRegisterJobsEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := s.RegisterJobs(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedJobCollection(res, "tiny")
		return vres, nil
	}
}

// NewUploadImageEndpoint returns an endpoint function that calls the method
// "uploadImage" of service "artworks".
func NewUploadImageEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UploadImagePayload)
		res, err := s.UploadImage(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedImage(res, "default")
		return vres, nil
	}
}
