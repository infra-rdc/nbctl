package netbox

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/client"
)

// newNetboxClient - The function creates a new NetBox client with the specified host, token, and HTTP scheme.
func newNetboxClient(netboxHost string, token string, httpScheme string) *client.NetBoxAPI {
	transport := httptransport.New(netboxHost, client.DefaultBasePath, []string{httpScheme})
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Token "+token)

	return client.New(transport, strfmt.Default)
}
