package query

// This File is modify from https://github.com/aws/aws-sdk-go/blob/main/private/protocol/query/unmarshal.go

import (
	"encoding/xml"

	"github.com/matpool/volcengine-go-sdk/private/protocol/xml/xmlutil"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineerr"
)

// UnmarshalHandler is a named request handler for unmarshaling volcenginequery protocol requests
var UnmarshalHandler = request.NamedHandler{Name: "awssdk.volcenginequery.Unmarshal", Fn: Unmarshal}

// UnmarshalMetaHandler is a named request handler for unmarshaling volcenginequery protocol request metadata
var UnmarshalMetaHandler = request.NamedHandler{Name: "awssdk.volcenginequery.UnmarshalMeta", Fn: UnmarshalMeta}

// Unmarshal unmarshals a response for an VOLCSTACK Query service.
func Unmarshal(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	if r.DataFilled() {
		decoder := xml.NewDecoder(r.HTTPResponse.Body)
		err := xmlutil.UnmarshalXML(r.Data, decoder, r.Operation.Name+"Result")
		if err != nil {
			r.Error = volcengineerr.NewRequestFailure(
				volcengineerr.New(request.ErrCodeSerialization, "failed decoding Query response", err),
				r.HTTPResponse.StatusCode,
				r.RequestID,
			)
			return
		}
	}
}

// UnmarshalMeta unmarshals header response values for an VOLCSTACK Query service.
func UnmarshalMeta(r *request.Request) {
	r.RequestID = r.HTTPResponse.Header.Get("X-Top-Requestid")
}
