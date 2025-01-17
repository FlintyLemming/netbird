//go:build !linux || android

package firewall

import (
	"context"
	"fmt"
	"runtime"

	log "github.com/sirupsen/logrus"

	firewall "github.com/FlintyLemming/netbird/client/firewall/manager"
	"github.com/FlintyLemming/netbird/client/firewall/uspfilter"
)

// NewFirewall creates a firewall manager instance
func NewFirewall(context context.Context, iface IFaceMapper) (firewall.Manager, error) {
	if !iface.IsUserspaceBind() {
		return nil, fmt.Errorf("not implemented for this OS: %s", runtime.GOOS)
	}

	// use userspace packet filtering firewall
	fm, err := uspfilter.Create(iface)
	if err != nil {
		return nil, err
	}
	err = fm.AllowNetbird()
	if err != nil {
		log.Warnf("failed to allow netbird interface traffic: %v", err)
	}
	return fm, nil
}
