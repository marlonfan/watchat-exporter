package driver

import (
	"errors"
	"fmt"
	"time"

	"github.com/marlonfan/watchat-exporter/metrics"
	"gorm.io/gorm"

	"github.com/marlonfan/watchat-exporter/config"
	"github.com/marlonfan/watchat-exporter/model"
	"github.com/mitchellh/mapstructure"
	"github.com/prometheus/client_golang/prometheus"
)

var ErrInvalidMetricType = errors.New("invalid metric type")

type mysqlDriver struct {
	name      string
	gauge     prometheus.Gauge
	count     prometheus.Counter
	cfg       *config.DriverConfig
	driverCfg *model.MysqlDriver
}

func newMysqlDriver(cfg *config.DriverConfig) (driver *mysqlDriver, err error) {
	c := &model.MysqlDriver{}
	if err := mapstructure.Decode(cfg.Config, &c); err != nil {
		return nil, err
	}
	err = validate.Struct(c)
	if err != nil {
		return nil, err
	}

	d := &mysqlDriver{
		cfg:       cfg,
		driverCfg: c,
	}

	switch cfg.MetricType {
	case "gauge":
		d.gauge = metrics.NewGauge(cfg.Name, cfg.Namespace, cfg.Subsystem, cfg.Labels)
	case "counter":
		d.count = metrics.NewCounter(cfg.Name, cfg.Namespace, cfg.Subsystem, cfg.Labels)
	default:
		return nil, ErrInvalidMetricType
	}

	d.name = generateName(cfg)
	return d, nil
}

func (m *mysqlDriver) Execute(source interface{}) error {
	switch m.cfg.MetricType {
	case "gauge":
		return m.gauger(source)
	case "count":
		return m.counter(source)
	default:
		return ErrInvalidMetricType
	}
}

func (m *mysqlDriver) gauger(source interface{}) error {
	var rst int64
	db, ok := source.(*gorm.DB)
	if !ok {
		return fmt.Errorf("source is not gorm.DB")
	}

	if err := db.Raw(m.driverCfg.Query, m.driverCfg.BindVar...).Scan(&rst).Error; err != nil {
		return err
	}
	m.gauge.Set(float64(rst))
	return nil
}

func (m *mysqlDriver) counter(source interface{}) error {
	var rst int64
	db, ok := source.(*gorm.DB)
	if !ok {
		return fmt.Errorf("source is not gorm.DB")
	}

	if err := db.Raw(m.driverCfg.Query, m.driverCfg.BindVar...).Scan(&rst).Error; err != nil {
		return err
	}

	m.count.Add(float64(rst))
	return nil
}

func (m *mysqlDriver) Name() string {
	return m.name
}

func (m *mysqlDriver) Source() string {
	return m.driverCfg.Source
}

func (m *mysqlDriver) Duration() time.Duration {
	d, err := time.ParseDuration(m.driverCfg.Interval)
	if err != nil {
		return time.Second * 5
	}
	return d
}
