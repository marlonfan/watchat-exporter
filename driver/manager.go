package driver

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/marlonfan/watchat-exporter/config"
	"github.com/marlonfan/watchat-exporter/constant"
)

var (
	// ErrDriverNotFound is returned when a driver is not found
	ErrDriverNotFound = errors.New("unknown driver")
	// ErrDriverConfigIsNil is returned when a driver config is nil
	ErrDriverConfigIsNil = errors.New("driver config is nil")
	// ErrDriverNotSupported is returned when a driver is not supported
	ErrDriverNotSupported = errors.New("driver not supported")
)

var validate = validator.New()

// Manager manages all the drivers
type Manager struct {
	drivers map[string]Driver
}

func newDriver(cfg *config.DriverConfig) (d Driver, err error) {
	if cfg == nil {
		return nil, ErrDriverConfigIsNil
	}

	switch cfg.Type {
	case constant.DriverMysql:
		return newMysqlDriver(cfg)
	default:
		return nil, ErrDriverNotSupported
	}
}

// NewManager creates a new driver manager
func NewManager(cfg *config.Config) (*Manager, error) {
	m := &Manager{
		drivers: make(map[string]Driver),
	}

	for _, driver := range cfg.Drivers {
		d, err := newDriver(driver)
		if err != nil {
			return nil, err
		}
		m.drivers[d.Name()] = d
	}

	return m, nil
}

// GetDriver returns a driver by name
func (m *Manager) GetDriver(name string) (Driver, error) {
	if d, ok := m.drivers[name]; ok {
		return d, nil
	}
	return nil, ErrDriverNotFound
}

// AddDriver adds a new driver
func (m *Manager) AddDriver(cfg *config.DriverConfig) (Driver, error) {
	if cfg == nil {
		return nil, ErrDriverConfigIsNil
	}

	d, err := newDriver(cfg)
	if err != nil {
		return nil, err
	}
	m.drivers[d.Name()] = d
	return d, nil
}

func (m *Manager) RemoveDriver(cfg *config.DriverConfig) (name string, err error) {
	if cfg == nil {
		return "", ErrDriverConfigIsNil
	}
	name = generateName(cfg)
	delete(m.drivers, name)
	return name, nil
}

// GetDrivers returns all the drivers
func (m *Manager) GetDrivers() map[string]Driver {
	return m.drivers
}
