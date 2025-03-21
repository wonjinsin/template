// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// PostUsersJSONBody defines parameters for PostUsers.
type PostUsersJSONBody struct {
	// Email Email of the user
	Email *string `json:"email,omitempty"`

	// Name Name of the user
	Name *string `json:"name,omitempty"`
}

// PutUsersIdJSONBody defines parameters for PutUsersId.
type PutUsersIdJSONBody struct {
	// Name Name of the user
	Name *string `json:"name,omitempty"`
}

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody PostUsersJSONBody

// PutUsersIdJSONRequestBody defines body for PutUsersId for application/json ContentType.
type PutUsersIdJSONRequestBody PutUsersIdJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetUsers request
	GetUsers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostUsersWithBody request with any body
	PostUsersWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostUsers(ctx context.Context, body PostUsersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteUsersId request
	DeleteUsersId(ctx context.Context, id uint64, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetUsersId request
	GetUsersId(ctx context.Context, id uint64, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutUsersIdWithBody request with any body
	PutUsersIdWithBody(ctx context.Context, id uint64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutUsersId(ctx context.Context, id uint64, body PutUsersIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetUsers(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUsersRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostUsersWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostUsersRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostUsers(ctx context.Context, body PostUsersJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostUsersRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteUsersId(ctx context.Context, id uint64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteUsersIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUsersId(ctx context.Context, id uint64, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUsersIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutUsersIdWithBody(ctx context.Context, id uint64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutUsersIdRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutUsersId(ctx context.Context, id uint64, body PutUsersIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutUsersIdRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetUsersRequest generates requests for GetUsers
func NewGetUsersRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostUsersRequest calls the generic PostUsers builder with application/json body
func NewPostUsersRequest(server string, body PostUsersJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostUsersRequestWithBody(server, "application/json", bodyReader)
}

// NewPostUsersRequestWithBody generates requests for PostUsers with any type of body
func NewPostUsersRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteUsersIdRequest generates requests for DeleteUsersId
func NewDeleteUsersIdRequest(server string, id uint64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetUsersIdRequest generates requests for GetUsersId
func NewGetUsersIdRequest(server string, id uint64) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPutUsersIdRequest calls the generic PutUsersId builder with application/json body
func NewPutUsersIdRequest(server string, id uint64, body PutUsersIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutUsersIdRequestWithBody(server, id, "application/json", bodyReader)
}

// NewPutUsersIdRequestWithBody generates requests for PutUsersId with any type of body
func NewPutUsersIdRequestWithBody(server string, id uint64, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetUsersWithResponse request
	GetUsersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUsersResponse, error)

	// PostUsersWithBodyWithResponse request with any body
	PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostUsersResponse, error)

	PostUsersWithResponse(ctx context.Context, body PostUsersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostUsersResponse, error)

	// DeleteUsersIdWithResponse request
	DeleteUsersIdWithResponse(ctx context.Context, id uint64, reqEditors ...RequestEditorFn) (*DeleteUsersIdResponse, error)

	// GetUsersIdWithResponse request
	GetUsersIdWithResponse(ctx context.Context, id uint64, reqEditors ...RequestEditorFn) (*GetUsersIdResponse, error)

	// PutUsersIdWithBodyWithResponse request with any body
	PutUsersIdWithBodyWithResponse(ctx context.Context, id uint64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutUsersIdResponse, error)

	PutUsersIdWithResponse(ctx context.Context, id uint64, body PutUsersIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutUsersIdResponse, error)
}

type GetUsersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// ResultCode Result Code
		ResultCode *string `json:"resultCode,omitempty"`
		ResultData *[]struct {
			// Email User Email
			Email *string `json:"email,omitempty"`

			// Id User ID
			Id *string `json:"id,omitempty"`

			// Name User Name
			Name *string `json:"name,omitempty"`
		} `json:"resultData,omitempty"`

		// ResultMsg Result Message
		ResultMsg *string `json:"resultMsg,omitempty"`

		// Trid Transaction ID
		Trid *string `json:"trid,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetUsersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUsersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostUsersResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// ResultCode Result Code
		ResultCode *string `json:"resultCode,omitempty"`
		ResultData *struct {
			// Email User Email
			Email *string `json:"email,omitempty"`

			// Id User ID
			Id *string `json:"id,omitempty"`

			// Name User Name
			Name *string `json:"name,omitempty"`
		} `json:"resultData,omitempty"`

		// ResultMsg Result Message
		ResultMsg *string `json:"resultMsg,omitempty"`

		// Trid Transaction ID
		Trid *string `json:"trid,omitempty"`
	}
	JSON400 *struct {
		// ResultCode Error code
		ResultCode *string `json:"resultCode,omitempty"`

		// ResultMsg Error message
		ResultMsg *string `json:"resultMsg,omitempty"`

		// Trid Transaction ID
		Trid *string `json:"trid,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r PostUsersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostUsersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteUsersIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// ResultCode Result Code
		ResultCode *string `json:"resultCode,omitempty"`

		// ResultMsg Result message
		ResultMsg *string `json:"resultMsg,omitempty"`

		// Trid Transaction ID
		Trid *string `json:"trid,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r DeleteUsersIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteUsersIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUsersIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// ResultCode Result Code
		ResultCode *string `json:"resultCode,omitempty"`
		ResultData *struct {
			// Email User Email
			Email *string `json:"email,omitempty"`

			// Id User ID
			Id *string `json:"id,omitempty"`

			// Name User Name
			Name *string `json:"name,omitempty"`
		} `json:"resultData,omitempty"`

		// ResultMsg Result Message
		ResultMsg *string `json:"resultMsg,omitempty"`

		// Trid Transaction ID
		Trid *string `json:"trid,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetUsersIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUsersIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PutUsersIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// ResultCode Result Code
		ResultCode *string `json:"resultCode,omitempty"`
		ResultData *struct {
			// Email User Email
			Email *string `json:"email,omitempty"`

			// Id User ID
			Id *string `json:"id,omitempty"`

			// Name User Name
			Name *string `json:"name,omitempty"`
		} `json:"resultData,omitempty"`

		// ResultMsg Result Message
		ResultMsg *string `json:"resultMsg,omitempty"`

		// Trid Transaction ID
		Trid *string `json:"trid,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r PutUsersIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PutUsersIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetUsersWithResponse request returning *GetUsersResponse
func (c *ClientWithResponses) GetUsersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUsersResponse, error) {
	rsp, err := c.GetUsers(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUsersResponse(rsp)
}

// PostUsersWithBodyWithResponse request with arbitrary body returning *PostUsersResponse
func (c *ClientWithResponses) PostUsersWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostUsersResponse, error) {
	rsp, err := c.PostUsersWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostUsersResponse(rsp)
}

func (c *ClientWithResponses) PostUsersWithResponse(ctx context.Context, body PostUsersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostUsersResponse, error) {
	rsp, err := c.PostUsers(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostUsersResponse(rsp)
}

// DeleteUsersIdWithResponse request returning *DeleteUsersIdResponse
func (c *ClientWithResponses) DeleteUsersIdWithResponse(ctx context.Context, id uint64, reqEditors ...RequestEditorFn) (*DeleteUsersIdResponse, error) {
	rsp, err := c.DeleteUsersId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteUsersIdResponse(rsp)
}

// GetUsersIdWithResponse request returning *GetUsersIdResponse
func (c *ClientWithResponses) GetUsersIdWithResponse(ctx context.Context, id uint64, reqEditors ...RequestEditorFn) (*GetUsersIdResponse, error) {
	rsp, err := c.GetUsersId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUsersIdResponse(rsp)
}

// PutUsersIdWithBodyWithResponse request with arbitrary body returning *PutUsersIdResponse
func (c *ClientWithResponses) PutUsersIdWithBodyWithResponse(ctx context.Context, id uint64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutUsersIdResponse, error) {
	rsp, err := c.PutUsersIdWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutUsersIdResponse(rsp)
}

func (c *ClientWithResponses) PutUsersIdWithResponse(ctx context.Context, id uint64, body PutUsersIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutUsersIdResponse, error) {
	rsp, err := c.PutUsersId(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutUsersIdResponse(rsp)
}

// ParseGetUsersResponse parses an HTTP response from a GetUsersWithResponse call
func ParseGetUsersResponse(rsp *http.Response) (*GetUsersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUsersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// ResultCode Result Code
			ResultCode *string `json:"resultCode,omitempty"`
			ResultData *[]struct {
				// Email User Email
				Email *string `json:"email,omitempty"`

				// Id User ID
				Id *string `json:"id,omitempty"`

				// Name User Name
				Name *string `json:"name,omitempty"`
			} `json:"resultData,omitempty"`

			// ResultMsg Result Message
			ResultMsg *string `json:"resultMsg,omitempty"`

			// Trid Transaction ID
			Trid *string `json:"trid,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePostUsersResponse parses an HTTP response from a PostUsersWithResponse call
func ParsePostUsersResponse(rsp *http.Response) (*PostUsersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostUsersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// ResultCode Result Code
			ResultCode *string `json:"resultCode,omitempty"`
			ResultData *struct {
				// Email User Email
				Email *string `json:"email,omitempty"`

				// Id User ID
				Id *string `json:"id,omitempty"`

				// Name User Name
				Name *string `json:"name,omitempty"`
			} `json:"resultData,omitempty"`

			// ResultMsg Result Message
			ResultMsg *string `json:"resultMsg,omitempty"`

			// Trid Transaction ID
			Trid *string `json:"trid,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest struct {
			// ResultCode Error code
			ResultCode *string `json:"resultCode,omitempty"`

			// ResultMsg Error message
			ResultMsg *string `json:"resultMsg,omitempty"`

			// Trid Transaction ID
			Trid *string `json:"trid,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseDeleteUsersIdResponse parses an HTTP response from a DeleteUsersIdWithResponse call
func ParseDeleteUsersIdResponse(rsp *http.Response) (*DeleteUsersIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteUsersIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// ResultCode Result Code
			ResultCode *string `json:"resultCode,omitempty"`

			// ResultMsg Result message
			ResultMsg *string `json:"resultMsg,omitempty"`

			// Trid Transaction ID
			Trid *string `json:"trid,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetUsersIdResponse parses an HTTP response from a GetUsersIdWithResponse call
func ParseGetUsersIdResponse(rsp *http.Response) (*GetUsersIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUsersIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// ResultCode Result Code
			ResultCode *string `json:"resultCode,omitempty"`
			ResultData *struct {
				// Email User Email
				Email *string `json:"email,omitempty"`

				// Id User ID
				Id *string `json:"id,omitempty"`

				// Name User Name
				Name *string `json:"name,omitempty"`
			} `json:"resultData,omitempty"`

			// ResultMsg Result Message
			ResultMsg *string `json:"resultMsg,omitempty"`

			// Trid Transaction ID
			Trid *string `json:"trid,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePutUsersIdResponse parses an HTTP response from a PutUsersIdWithResponse call
func ParsePutUsersIdResponse(rsp *http.Response) (*PutUsersIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutUsersIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// ResultCode Result Code
			ResultCode *string `json:"resultCode,omitempty"`
			ResultData *struct {
				// Email User Email
				Email *string `json:"email,omitempty"`

				// Id User ID
				Id *string `json:"id,omitempty"`

				// Name User Name
				Name *string `json:"name,omitempty"`
			} `json:"resultData,omitempty"`

			// ResultMsg Result Message
			ResultMsg *string `json:"resultMsg,omitempty"`

			// Trid Transaction ID
			Trid *string `json:"trid,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get users
	// (GET /users)
	GetUsers(ctx echo.Context) error
	// Create user
	// (POST /users)
	PostUsers(ctx echo.Context) error
	// Delete user
	// (DELETE /users/{id})
	DeleteUsersId(ctx echo.Context, id uint64) error
	// Get user by ID
	// (GET /users/{id})
	GetUsersId(ctx echo.Context, id uint64) error
	// Update user
	// (PUT /users/{id})
	PutUsersId(ctx echo.Context, id uint64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsers(ctx)
	return err
}

// PostUsers converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUsers(ctx)
	return err
}

// DeleteUsersId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUsersId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint64

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUsersId(ctx, id)
	return err
}

// GetUsersId converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint64

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsersId(ctx, id)
	return err
}

// PutUsersId converts echo context to params.
func (w *ServerInterfaceWrapper) PutUsersId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint64

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PutUsersId(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/users", wrapper.GetUsers)
	router.POST(baseURL+"/users", wrapper.PostUsers)
	router.DELETE(baseURL+"/users/:id", wrapper.DeleteUsersId)
	router.GET(baseURL+"/users/:id", wrapper.GetUsersId)
	router.PUT(baseURL+"/users/:id", wrapper.PutUsersId)

}
