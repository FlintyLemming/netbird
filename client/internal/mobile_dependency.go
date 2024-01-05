package internal

import (
	"github.com/FlintyLemming/netbird/client/internal/dns"
	"github.com/FlintyLemming/netbird/client/internal/listener"
	"github.com/FlintyLemming/netbird/client/internal/stdnet"
	"github.com/FlintyLemming/netbird/iface"
)

// MobileDependency collect all dependencies for mobile platform
type MobileDependency struct {
	// Android only
	TunAdapter            iface.TunAdapter
	IFaceDiscover         stdnet.ExternalIFaceDiscover
	NetworkChangeListener listener.NetworkChangeListener
	HostDNSAddresses      []string
	DnsReadyListener      dns.ReadyListener

	//	iOS only
	DnsManager     dns.IosDnsManager
	FileDescriptor int32
}
