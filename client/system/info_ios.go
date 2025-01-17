//go:build ios
// +build ios

package system

import (
	"context"
	"runtime"

	"github.com/FlintyLemming/netbird/version"
)

// GetInfo retrieves and parses the system information
func GetInfo(ctx context.Context) *Info {

	// Convert fixed-size byte arrays to Go strings
	sysName := extractOsName(ctx, "sysName")
	swVersion := extractOsVersion(ctx, "swVersion")

	gio := &Info{Kernel: sysName, OSVersion: swVersion, Core: swVersion, Platform: "unknown", OS: sysName, GoOS: runtime.GOOS, CPUs: runtime.NumCPU()}
	gio.Hostname = extractDeviceName(ctx, "hostname")
	gio.WiretrusteeVersion = version.NetbirdVersion()
	gio.UIVersion = extractUserAgent(ctx)

	return gio
}

// extractOsVersion extracts operating system version from context or returns the default
func extractOsVersion(ctx context.Context, defaultName string) string {
	v, ok := ctx.Value(OsVersionCtxKey).(string)
	if !ok {
		return defaultName
	}
	return v
}

// extractOsName extracts operating system name from context or returns the default
func extractOsName(ctx context.Context, defaultName string) string {
	v, ok := ctx.Value(OsNameCtxKey).(string)
	if !ok {
		return defaultName
	}
	return v
}
