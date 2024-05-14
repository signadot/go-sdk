package transport

import (
	"fmt"
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
	APIURL          string
	ArtifactsAPIURL string
	UserAgent       string
	Debug           bool
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
	rt := oaclient.New(tc.Host, tc.BasePath, tc.Schemes)

	// setup auth
	if conf.APIKey != "" {
		rt.DefaultAuthentication = oaclient.APIKeyAuth(APIKeyHeader, "header", conf.APIKey)
	} else if conf.ClusterToken != "" {
		rt.DefaultAuthentication = oaclient.APIKeyAuth(ClusterTokenHeader, "header", conf.ClusterToken)
	}

	// enable debug
	rt.SetDebug(conf.Debug)

	// add User-Agent to every request
	rt.Transport = &userAgent{
		inner: rt.Transport,
		agent: conf.UserAgent,
	}

	return FixAPIErrors(rt)
}
