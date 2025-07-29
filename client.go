package depscian_client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	apiclient "go.depscian.tech/internal/client"
)

const (
	defaultBaseURL = "https://api.depscian.tech/v2"
	apiKeyHeader   = "X-API-Key"
)

type service struct {
	client *apiclient.ClientWithResponses
}

type Client struct {
	client *apiclient.ClientWithResponses
	common service

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

	client, err := apiclient.NewClientWithResponses(
		baseURL,
		apiclient.WithHTTPClient(httpClient),
		apiclient.WithRequestEditorFn(authInterceptor),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create API client: %w", err)
	}

	c := &Client{client: client}
	c.common.client = client
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

func (s *AdminsService) Get(ctx context.Context, serverID int) (*apiclient.GetServerAdminsV2AdminsGetResponse, error) {
	params := &apiclient.GetServerAdminsV2AdminsGetParams{ServerId: serverID}
	return s.client.GetServerAdminsV2AdminsGetWithResponse(ctx, params)
}

type FamiliesService service

func (s *FamiliesService) List(ctx context.Context, serverID int) (*apiclient.GetFamiliesV2FamiliesGetResponse, error) {
	params := &apiclient.GetFamiliesV2FamiliesGetParams{ServerId: serverID}
	return s.client.GetFamiliesV2FamiliesGetWithResponse(ctx, params)
}

func (s *FamiliesService) Get(ctx context.Context, serverID, famID int) (*apiclient.GetFamilyV2FamilyGetResponse, error) {
	params := &apiclient.GetFamilyV2FamilyGetParams{ServerId: serverID, FamId: famID}
	return s.client.GetFamilyV2FamilyGetWithResponse(ctx, params)
}

type FractionsService service

func (s *FractionsService) List(ctx context.Context, serverID int) (*apiclient.GetFractionsListV2FractionsGetResponse, error) {
	params := &apiclient.GetFractionsListV2FractionsGetParams{ServerId: serverID}
	return s.client.GetFractionsListV2FractionsGetWithResponse(ctx, params)
}

func (s *FractionsService) GetMembers(ctx context.Context, serverID int, fractionID string) (*apiclient.GetFractionMembersV2FractionGetResponse, error) {
	params := &apiclient.GetFractionMembersV2FractionGetParams{ServerId: serverID, FractionId: fractionID}
	return s.client.GetFractionMembersV2FractionGetWithResponse(ctx, params)
}

type GhettoService service

func (s *GhettoService) Get(ctx context.Context, serverID int) (*apiclient.GetGhettoListV2GhettoGetResponse, error) {
	params := &apiclient.GetGhettoListV2GhettoGetParams{ServerId: serverID}
	return s.client.GetGhettoListV2GhettoGetWithResponse(ctx, params)
}

type LeadershipService service

func (s *LeadershipService) GetLeaders(ctx context.Context, serverID int) (*apiclient.GetLeadersListV2LeadersGetResponse, error) {
	params := &apiclient.GetLeadersListV2LeadersGetParams{ServerId: serverID}
	return s.client.GetLeadersListV2LeadersGetWithResponse(ctx, params)
}

func (s *LeadershipService) GetSubleaders(ctx context.Context, serverID int) (*apiclient.GetSubleadersListV2SubleadersGetResponse, error) {
	params := &apiclient.GetSubleadersListV2SubleadersGetParams{ServerId: serverID}
	return s.client.GetSubleadersListV2SubleadersGetWithResponse(ctx, params)
}

type MapService service

func (s *MapService) Get(ctx context.Context, serverID int) (*apiclient.GetPropertyMapWithPoiV2MapGetResponse, error) {
	params := &apiclient.GetPropertyMapWithPoiV2MapGetParams{ServerId: serverID}
	return s.client.GetPropertyMapWithPoiV2MapGetWithResponse(ctx, params)
}

type OnlineService service

func (s *OnlineService) Get(ctx context.Context, serverID int) (*apiclient.GetOnlineListV2OnlineGetResponse, error) {
	params := &apiclient.GetOnlineListV2OnlineGetParams{ServerId: serverID}
	return s.client.GetOnlineListV2OnlineGetWithResponse(ctx, params)
}

type PlayerService service

func (s *PlayerService) Find(ctx context.Context, serverID int, nickname string) (*apiclient.FindPlayerV2PlayerFindGetResponse, error) {
	params := &apiclient.FindPlayerV2PlayerFindGetParams{ServerId: serverID, Nickname: nickname}
	return s.client.FindPlayerV2PlayerFindGetWithResponse(ctx, params)
}

type SobesService service

func (s *SobesService) Get(ctx context.Context, serverID int) (*apiclient.GetSobesListV2SobesGetResponse, error) {
	params := &apiclient.GetSobesListV2SobesGetParams{ServerId: serverID}
	return s.client.GetSobesListV2SobesGetWithResponse(ctx, params)
}

type StatusService service

func (s *StatusService) Get(ctx context.Context) (*apiclient.GetStatusV2StatusGetResponse, error) {
	return s.client.GetStatusV2StatusGetWithResponse(ctx)
}
