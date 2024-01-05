//go:build android

package routemanager

import (
	"context"
	"fmt"

	firewall "github.com/FlintyLemming/netbird/client/firewall/manager"
	"github.com/FlintyLemming/netbird/iface"
)

func newServerRouter(context.Context, *iface.WGIface, firewall.Manager) (serverRouter, error) {
	return nil, fmt.Errorf("server route not supported on this os")
}
