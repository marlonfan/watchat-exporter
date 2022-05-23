package driver

import (
	"fmt"
	"time"

	"github.com/marlonfan/watchat-exporter/config"
)

// Driver is the interface that must be implemented by a driver.
type Driver interface {
	Name() string
	Source() string
	Duration() time.Duration
	Execute(source interface{}) (err error)
}

// generateName generates a name for a driver.
func generateName(cfg *config.DriverConfig) string {
	labelStr := ""
	for k, v := range cfg.Labels {
		labelStr += ":" + k + "=" + v
	}
	return fmt.Sprintf("%s%s", cfg.Name, labelStr)
}
