package cstrace

import (
	"go.opentelemetry.io/otel/api/key"
	"go.opentelemetry.io/otel/api/trace"
	"google.golang.org/grpc/codes"
)

var ErrorKey = key.New("error")

// Status codes for use with Span.SetStatus. These correspond to the status
// codes used by gRPC defined here: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto

// Status if there is an error, it sets the error code "unknown" with the error
// string as span status otherwise status ok.
func Status(span trace.Span, err error) {
	if err == nil {
		span.SetStatus(codes.OK)
		return
	}
	span.SetStatus(codes.Unknown)
	span.SetAttributes(ErrorKey.String(err.Error()))
}

// StatusErrorWithCode sets a custom code with an error.
// go.opencensus.io/trace/status_codes.go. These correspond to the status codes
// used by gRPC defined here:
// https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
