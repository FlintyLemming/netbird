package NetBirdSDK

import (
	"github.com/FlintyLemming/netbird/util"
)

// InitializeLog initializes the log file.
func InitializeLog(logLevel string, filePath string) error {
	return util.InitLog(logLevel, filePath)
}
