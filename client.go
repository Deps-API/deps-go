package depsclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	apiclient "go.depscian.tech/internal/client"
)

var ErrNotFound = errors.New("not found")

const (
	defaultBaseURL = "https://api.depscian.tech/v2"
	apiKeyHeader   = "X-API-Key"
)

func unpackResponse[T any](
	responseBody *T,
	httpResponse *http.Response,
	err error,
) (*T, error) {
	if err != nil {
		return nil, fmt.Errorf("client execution error: %w", err)
	}

	if httpResponse.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}

	if httpResponse.StatusCode < 200 || httpResponse.StatusCode >= 300 {
		return nil, fmt.Errorf("api error: status %s", httpResponse.Status)
	}

	if responseBody == nil {
		return nil, ErrNotFound
	}

	return responseBody, nil
}

type service struct {
	client *Client
}

type Client struct {
	internalClient *apiclient.ClientWithResponses
	common         service

	Admins     *AdminsService
	Families   *FamiliesService
	Fractions  *FractionsService
	Ghetto     *GhettoService
	Leadership *LeadershipService
	Map        *MapService
	Online     *OnlineService
	Player     *PlayerService
	Sobes      *SobesService
	Status     *StatusService
}

type Option func(*http.Client, *string) error

func NewClient(apiKey string, opts ...Option) (*Client, error) {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	baseURL := defaultBaseURL

	for _, opt := range opts {
		if err := opt(httpClient, &baseURL); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	authInterceptor := func(ctx context.Context, req *http.Request) error {
		req.Header.Set(apiKeyHeader, apiKey)
		return nil
	}

	generatedClient, err := apiclient.NewClientWithResponses(
		baseURL,
		apiclient.WithHTTPClient(httpClient),
		apiclient.WithRequestEditorFn(authInterceptor),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create API client: %w", err)
	}

	c := &Client{internalClient: generatedClient}
	c.common.client = c
	c.Admins = (*AdminsService)(&c.common)
	c.Families = (*FamiliesService)(&c.common)
	c.Fractions = (*FractionsService)(&c.common)
	c.Ghetto = (*GhettoService)(&c.common)
	c.Leadership = (*LeadershipService)(&c.common)
	c.Map = (*MapService)(&c.common)
	c.Online = (*OnlineService)(&c.common)
	c.Player = (*PlayerService)(&c.common)
	c.Sobes = (*SobesService)(&c.common)
	c.Status = (*StatusService)(&c.common)

	return c, nil
}

func WithBaseURL(url string) Option {
	return func(_ *http.Client, baseURL *string) error {
		*baseURL = url
		return nil
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *http.Client, _ *string) error {
		c.Timeout = timeout
		return nil
	}
}

func WithHTTPClient(customClient *http.Client) Option {
	return func(c *http.Client, _ *string) error {
		*c = *customClient
		return nil
	}
}

type AdminsService service

func (s *AdminsService) Get(ctx context.Context, serverID int) (*apiclient.AdminsResponse, error) {
	params := &apiclient.GetServerAdminsV2AdminsGetParams{ServerId: serverID}
	resp, err := s.client.internalClient.GetServerAdminsV2AdminsGetWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return unpackResponse(resp.JSON200, resp.HTTPResponse, err)
}

type FamiliesService service

func (s *FamiliesService) List(ctx context.Context, serverID int) (*apiclient.FamilyListResponse, error) {
	params := &apiclient.GetFamiliesV2FamiliesGetParams{ServerId: serverID}
	resp, err := s.client.internalClient.GetFamiliesV2FamiliesGetWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return unpackResponse(resp.JSON200, resp.HTTPResponse, err)
}

func (s *FamiliesService) Get(ctx context.Context, serverID, famID int) (*apiclient.FamilyResponse, error) {
	params := &apiclient.GetFamilyV2FamilyGetParams{ServerId: serverID, FamId: famID}
	resp, err := s.client.internalClient.GetFamilyV2FamilyGetWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return unpackResponse(resp.JSON200, resp.HTTPResponse, err)
}

type FractionsService service

func (s *FractionsService) List(ctx context.Context, serverID int) (*apiclient.FractionsListResponse, error) {
	params := &apiclient.GetFractionsListV2FractionsGetParams{ServerId: serverID}
	resp, err := s.client.internalClient.GetFractionsListV2FractionsGetWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return unpackResponse(resp.JSON200, resp.HTTPResponse, err)
}

func (s *FractionsService) GetMembers(ctx context.Context, serverID int, fractionID string) (*apiclient.FractionResponse, error) {
	params := &apiclient.GetFractionMembersV2FractionGetParams{ServerId: serverID, FractionId: fractionID}
	resp, err := s.client.internalClient.GetFractionMembersV2FractionGetWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return unpackResponse(resp.JSON200, resp.HTTPResponse, err)
}

type GhettoService service

func (s *GhettoService) Get(ctx context.Context, serverID int) (*apiclient.GhettoResponse, error) {
	params := &apiclient.GetGhettoListV2GhettoGetParams{ServerId: serverID}
	resp, err := s.client.internalClient.GetGhettoListV2GhettoGetWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return unpackResponse(resp.JSON200, resp.HTTPResponse, err)
}

type LeadershipService service

func (s *LeadershipService) GetLeaders(ctx context.Context, serverID int) (*apiclient.LeadersResponse, error) {
	params := &apiclient.GetLeadersListV2LeadersGetParams{ServerId: serverID}
	resp, err := s.client.internalClient.GetLeadersListV2LeadersGetWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return unpackResponse(resp.JSON200, resp.HTTPResponse, err)
}

func (s *LeadershipService) GetSubleaders(ctx context.Context, serverID int) (*apiclient.SubleadersResponse, error) {
	params := &apiclient.GetSubleadersListV2SubleadersGetParams{ServerId: serverID}
	resp, err := s.client.internalClient.GetSubleadersListV2SubleadersGetWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return unpackResponse(resp.JSON200, resp.HTTPResponse, err)
}

type MapService service

func (s *MapService) Get(ctx context.Context, serverID int) (*apiclient.MapResponse, error) {
	params := &apiclient.GetPropertyMapWithPoiV2MapGetParams{ServerId: serverID}
	resp, err := s.client.internalClient.GetPropertyMapWithPoiV2MapGetWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return unpackResponse(resp.JSON200, resp.HTTPResponse, err)
}

type OnlineService service

func (s *OnlineService) Get(ctx context.Context, serverID int) (*apiclient.OnlinePlayersResponse, error) {
	params := &apiclient.GetOnlineListV2OnlineGetParams{ServerId: serverID}
	resp, err := s.client.internalClient.GetOnlineListV2OnlineGetWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return unpackResponse(resp.JSON200, resp.HTTPResponse, err)
}

type PlayerService service

func (s *PlayerService) Find(ctx context.Context, serverID int, nickname string) (*apiclient.PlayerResponse, error) {
	params := &apiclient.FindPlayerV2PlayerFindGetParams{ServerId: serverID, Nickname: nickname}
	resp, err := s.client.internalClient.FindPlayerV2PlayerFindGetWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return unpackResponse(resp.JSON200, resp.HTTPResponse, err)
}

type SobesService service

func (s *SobesService) Get(ctx context.Context, serverID int) (*apiclient.SobesResponse, error) {
	params := &apiclient.GetSobesListV2SobesGetParams{ServerId: serverID}
	resp, err := s.client.internalClient.GetSobesListV2SobesGetWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}
	return unpackResponse(resp.JSON200, resp.HTTPResponse, err)
}

type StatusService service

func (s *StatusService) Get(ctx context.Context) (*apiclient.StatusResponse, error) {
	resp, err := s.client.internalClient.GetStatusV2StatusGetWithResponse(ctx)
	if err != nil {
		return nil, err
	}
	return unpackResponse(resp.JSON200, resp.HTTPResponse, err)
}
