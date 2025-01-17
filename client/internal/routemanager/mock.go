package routemanager

import (
	"context"
	"fmt"

	firewall "github.com/FlintyLemming/netbird/client/firewall/manager"
	"github.com/FlintyLemming/netbird/client/internal/listener"
	"github.com/FlintyLemming/netbird/iface"
	"github.com/FlintyLemming/netbird/route"
)

// MockManager is the mock instance of a route manager
type MockManager struct {
	UpdateRoutesFunc func(updateSerial uint64, newRoutes []*route.Route) error
	StopFunc         func()
}

// InitialRouteRange mock implementation of InitialRouteRange from Manager interface
func (m *MockManager) InitialRouteRange() []string {
	return nil
}

// UpdateRoutes mock implementation of UpdateRoutes from Manager interface
func (m *MockManager) UpdateRoutes(updateSerial uint64, newRoutes []*route.Route) error {
	if m.UpdateRoutesFunc != nil {
		return m.UpdateRoutesFunc(updateSerial, newRoutes)
	}
	return fmt.Errorf("method UpdateRoutes is not implemented")
}

// Start mock implementation of Start from Manager interface
func (m *MockManager) Start(ctx context.Context, iface *iface.WGIface) {
}

// SetRouteChangeListener mock implementation of SetRouteChangeListener from Manager interface
func (m *MockManager) SetRouteChangeListener(listener listener.NetworkChangeListener) {

}

func (m *MockManager) EnableServerRouter(firewall firewall.Manager) error {
	panic("implement me")
}

// Stop mock implementation of Stop from Manager interface
func (m *MockManager) Stop() {
	if m.StopFunc != nil {
		m.StopFunc()
	}
}
