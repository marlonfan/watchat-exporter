package source

import (
	"fmt"

	"github.com/marlonfan/watchat-exporter/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager struct {
	sources map[string]interface{}
}

// NewManager creates a new Source Manager
func NewManager(cfg *config.Config) (*Manager, error) {
	m := &Manager{
		sources: make(map[string]interface{}),
	}
	for _, dcfg := range cfg.Sources {
		s, err := NewSource(dcfg)
		if err != nil {
			return nil, err
		}
		m.sources[dcfg.Name] = s
	}

	return m, nil
}

func (m *Manager) Get(name string) (interface{}, error) {
	s, ok := m.sources[name]
	if !ok {
		return nil, fmt.Errorf("source %s not found", name)
	}
	return s, nil
}

// Add a new source, replacing an existing one if it exists
func (m *Manager) Add(s *config.SourceConfig) error {
	source, err := NewSource(s)
	if err != nil {
		return err
	}
	m.sources[s.Name] = source
	return nil
}

// NewSource creates a new Source
func NewSource(cfg *config.SourceConfig) (source interface{}, err error) {
	switch cfg.Type {
	case "mysql":
		oDSN, ok := cfg.Config["dsn"]
		if !ok {
			return nil, fmt.Errorf("mysql dsn not found")
		}
		dsn, ok := oDSN.(string)
		if !ok {
			return nil, fmt.Errorf("mysql dsn not string")
		}
		return gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unknown source type: %s", cfg.Type)
	}
}
