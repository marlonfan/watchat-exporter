package watchat

import (
	"errors"
	"fmt"
	"time"

	"github.com/marlonfan/watchat-exporter/config"
	"github.com/marlonfan/watchat-exporter/driver"
	"github.com/marlonfan/watchat-exporter/source"
	"github.com/rfyiamcool/go-timewheel"
)

const (
	defaultTimeWheelTickUnit = time.Millisecond * 200
	defaultTimeWheelTickSize = 1 << 10
)

// Client watch a data source and execute a query command
type Client struct {
	driverManager      *driver.Manager
	querySourceManager *source.Manager
	timewheel          *timewheel.TimeWheel
	taskList           map[string]*timewheel.Task
}

// GetDriverManager get driver manager
func (w *Client) GetDriverManager() *driver.Manager {
	return w.driverManager
}

// GetQuerySourceManager get query source manager
func (w *Client) GetQuerySourceManager() *source.Manager {
	return w.querySourceManager
}

// GetTimeWheel get time wheel
func (w *Client) GetTimeWheel() interface{} {
	return w.timewheel
}

// NewWatchat creates a new watchat
func NewWatchat(cfg *config.Config) (*Client, error) {
	var err error
	w := &Client{}

	w.querySourceManager, err = source.NewManager(cfg)
	if err != nil {
		return nil, err
	}

	w.taskList = make(map[string]*timewheel.Task)
	w.timewheel, err = timewheel.NewTimeWheel(defaultTimeWheelTickUnit, defaultTimeWheelTickSize)
	if err != nil {
		return nil, err
	}

	w.driverManager, err = driver.NewManager(cfg)
	if err != nil {
		return nil, err
	}

	// init
	for _, d := range w.driverManager.GetDrivers() {
		// call driver to init, and add cron
		caller := w.createCaller(d)
		if !cfg.Lazy {
			caller()
		}
		w.taskList[d.Name()] = w.timewheel.AddCron(d.Duration(), caller)
	}

	return w, nil
}

func (w *Client) createCaller(d driver.Driver) func() {
	return func() {
		s, err := w.querySourceManager.Get(d.Source())
		if err != nil {
			fmt.Println(err)
			return
		}
		err = d.Execute(s)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

// Start start watchat
func (w *Client) Start() {
	w.timewheel.Start()
}

// AddSource creates a new source from config
func (w *Client) AddSource(s *config.SourceConfig) error {
	return w.querySourceManager.Add(s)
}

// AddDriver creates a new driver from config
func (w *Client) AddDriver(cfg *config.DriverConfig) (err error) {
	var d driver.Driver
	if d, _ = w.driverManager.GetDriver(cfg.Name); d != nil {
		return errors.New("driver already exists")
	}

	if d, err = w.driverManager.AddDriver(cfg); err != nil {
		return err
	}

	if _, err := w.querySourceManager.Get(d.Source()); err != nil {
		_,_ = w.driverManager.RemoveDriver(cfg)
		return err
	}

	caller := w.createCaller(d)
	caller()
	w.taskList[d.Name()] = w.timewheel.AddCron(d.Duration(), caller)
	return nil
}

func (w *Client) RemoveDriver(cfg *config.DriverConfig) (err error) {
	name, err := w.driverManager.RemoveDriver(cfg)
	if err != nil {
		return err
	}
	return w.timewheel.Remove(w.taskList[name])
}

// Stop stop watchat
func (w *Client) Stop() {
	w.timewheel.Stop()
}
