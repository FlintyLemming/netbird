//go:build !android

package test

import firewall "github.com/FlintyLemming/netbird/client/firewall/manager"

var (
	InsertRuleTestCases = []struct {
		Name      string
		InputPair firewall.RouterPair
	}{
		{
			Name: "Insert Forwarding IPV4 Rule",
			InputPair: firewall.RouterPair{
				ID:          "zxa",
				Source:      "100.100.100.1/32",
				Destination: "100.100.200.0/24",
				Masquerade:  false,
			},
		},
		{
			Name: "Insert Forwarding And Nat IPV4 Rules",
			InputPair: firewall.RouterPair{
				ID:          "zxa",
				Source:      "100.100.100.1/32",
				Destination: "100.100.200.0/24",
				Masquerade:  true,
			},
		},
	}

	RemoveRuleTestCases = []struct {
		Name      string
		InputPair firewall.RouterPair
		IpVersion string
	}{
		{
			Name: "Remove Forwarding And Nat IPV4 Rules",
			InputPair: firewall.RouterPair{
				ID:          "zxa",
				Source:      "100.100.100.1/32",
				Destination: "100.100.200.0/24",
				Masquerade:  true,
			},
		},
	}
)
