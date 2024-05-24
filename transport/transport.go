package transport

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-openapi/runtime"
	oaclient "github.com/go-openapi/runtime/client"
	"github.com/signadot/go-sdk/client"
)

const (
	APIKeyHeader       = "signadot-api-key"
	ClusterTokenHeader = "signadot-cluster-token"
)

type APIConfig struct {
	APIKey          string
	ClusterToken    string
	BearerToken     string
	APIURL          string
	ArtifactsAPIURL string
	UserAgent       string
	Debug           bool

	// Advanced (optional) settings
	Consumers  map[string]runtime.Consumer
	Producers  map[string]runtime.Producer
	HTTPClient *http.Client
}

func InitAPITransport(conf *APIConfig) (runtime.ClientTransport, error) {
	// create the base transport for accessing the base api
	baseTC := client.DefaultTransportConfig()
	if conf.APIURL != "" {
		u, err := url.Parse(conf.APIURL)
		if err != nil {
			return nil, fmt.Errorf("invalid api url: %w", err)
		}
		baseTC.Host = u.Host
		baseTC.Schemes = []string{u.Scheme}
	}

	baseT := createAPITransport(baseTC, conf)
	if conf.ArtifactsAPIURL == "" {
		// just return the base transport
		return baseT, nil
	}

	// create another transport for accessing the artifacts api
	u, err := url.Parse(conf.ArtifactsAPIURL)
	if err != nil {
		return nil, fmt.Errorf("invalid artifacts api url: %w", err)
	}
	artifactsTC := client.DefaultTransportConfig()
	artifactsTC.Host = u.Host
	artifactsTC.Schemes = []string{u.Scheme}
	artifactsT := createAPITransport(artifactsTC, conf)

	return apiComposableTransport{
		base:      baseT,
		artifacts: artifactsT,
	}, nil
}

func createAPITransport(tc *client.TransportConfig, conf *APIConfig) runtime.ClientTransport {
	rt := oaclient.NewWithClient(tc.Host, tc.BasePath, tc.Schemes, conf.HTTPClient)

	// setup auth
	switch {
	case conf.BearerToken != "":
		rt.DefaultAuthentication = oaclient.BearerToken(conf.BearerToken)
	case conf.APIKey != "":
		rt.DefaultAuthentication = oaclient.APIKeyAuth(APIKeyHeader, "header", conf.APIKey)
	case conf.ClusterToken != "":
		rt.DefaultAuthentication = oaclient.APIKeyAuth(ClusterTokenHeader, "header", conf.ClusterToken)
	}

	// enable debug
	rt.SetDebug(conf.Debug)

	// add User-Agent to every request
	rt.Transport = &userAgent{
		inner: rt.Transport,
		agent: conf.UserAgent,
	}

	// apply consumer and producer settings
	for mimeType, consumer := range conf.Consumers {
		rt.Consumers[mimeType] = consumer
	}
	for mimeType, producer := range conf.Producers {
		rt.Producers[mimeType] = producer
	}

	return FixAPIErrors(rt)
}
