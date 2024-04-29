package dap

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	libhttp "net/http"
	liburl "net/url"
	"time"

	"github.com/alecthomas/types/optional"
	"github.com/tbd54566975/web5-go/dids"
	"github.com/tbd54566975/web5-go/dids/didcore"
)

type Client struct {
	http *libhttp.Client
}

type newClientOptions struct {
	http *libhttp.Client
}

type NewClientOption func(*newClientOptions)

func HTTP(client *libhttp.Client) NewClientOption {
	return func(opts *newClientOptions) {
		opts.http = client
	}
}

func NewClient(opts ...NewClientOption) Client {
	options := newClientOptions{
		http: &libhttp.Client{
			Timeout: 15 * time.Second,
		},
	}

	for _, o := range opts {
		o(&options)
	}

	return Client{http: options.http}
}

func (c Client) Register(ctx context.Context, r RegistrationRequest) (*RegistrationResponse, error) {
	url, err := c.getRegistryURL(r.Domain)
	if err != nil {
		return nil, err
	}

	url.Path = "/daps"

	body, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	reader := bytes.NewReader(body)

	req, err := libhttp.NewRequestWithContext(ctx, http.MethodPost, url.String(), reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var responseBody HTTPResponse[RegistrationResponse]
	err = json.Unmarshal(responseBodyBytes, &responseBody)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		err := responseBody.Error.Default(ErrHTTPResponse{Message: resp.Status})
		err.Status = resp.StatusCode

		return nil, err
	}

	return responseBody.Data.Ptr(), nil
}

func (c Client) Resolve(ctx context.Context, input string) (*ResolutionResponse, error) {
	dap, err := Parse(input)
	if err != nil {
		return nil, fmt.Errorf("invalid dap: %w", err)
	}

	url, err := c.getRegistryURL(dap.Domain)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve domain did: %w", err)
	}

	url.Path = "/" + dap.String()

	req, err := libhttp.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var responseBody HTTPResponse[ResolutionResponse]
	err = json.Unmarshal(responseBodyBytes, &responseBody)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		err := responseBody.Error.Default(ErrHTTPResponse{Message: resp.Status})
		err.Status = resp.StatusCode

		return nil, err
	}

	return responseBody.Data.Ptr(), nil
}

func (c Client) getRegistryURL(domain string) (*liburl.URL, error) {
	hostDID := "did:web" + domain

	resolution, err := dids.Resolve(hostDID)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve %s: %w", hostDID, err)
	}

	var dapService *didcore.Service
	for _, s := range resolution.Document.Service {
		if s.Type == ServiceType {
			dapService = &s
			break
		}
	}

	if dapService == nil {
		return nil, fmt.Errorf("no %s service found for %s", ServiceType, hostDID)
	}

	registryURL, err := liburl.Parse(dapService.ServiceEndpoint[0])
	if err != nil {
		return nil, fmt.Errorf("invalid registry URL %s: %w", dapService.ServiceEndpoint[0], err)
	}

	return registryURL, nil
}

type HTTPResponse[T any] struct {
	Data  optional.Option[T]               `json:"data,omitempty"`
	Error optional.Option[ErrHTTPResponse] `json:"error,omitempty"`
}

type ErrHTTPResponse struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e ErrHTTPResponse) Error() string {
	return fmt.Sprintf("(%d) %s", e.Status, e.Message)
}
