package watchat

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/marlonfan/watchat-exporter/config"
	"gopkg.in/yaml.v3"
)

func TestNewWatchat(t *testing.T) {
	cfg := &config.Config{
		Drivers: []*config.DriverConfig{
			{
				Name: "mysqlTest",
				Labels: map[string]string{
					"group":   "arch",
					"service": "elance",
					"module":  "goods",
				},
				MetricType: "gauge",
				Help:       "test help",
				Type:       "watchat.driver.mysql",
				Config: map[string]interface{}{
					"query": "select count(*) from users where id > ?",
					"bind_var": []interface{}{
						0,
					},
					"interval": "3s",
					"source":   "mysqlTest",
				},
			},
		},
		Sources: []*config.SourceConfig{
			{
				Name: "mysqlTest",
				Type: "mysql",
				Config: map[string]interface{}{
					"dsn": "root:123456@@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
				},
			},
		},
	}
	w, err := NewWatchat(cfg)
	if err != nil {
		t.Error(err)
		return
	}
	if w == nil {
		t.Error("NewWatchat() should not return nil")
		return
	}
	w.Start()

	time.Sleep(time.Second * 10)
	w.Stop()
}

func TestWatchat_FromYAML(t *testing.T) {
	var cfg *config.Config
	b, err := ioutil.ReadFile("./testdata/config.yaml")
	if err != nil {
		t.Error(err)
	}
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		t.Error(err)
	}
	w, err := NewWatchat(cfg)
	if err != nil {
		t.Error(err)
		return
	}
	if w == nil {
		t.Error("NewWatchat() should not return nil")
		return
	}
	w.Start()

	time.Sleep(time.Second * 10)
	w.Stop()
}

func TestWatchat_AddDriver(t *testing.T) {
	cfg := &config.Config{}
	w, err := NewWatchat(cfg)
	if err != nil {
		t.Error(err)
		return
	}
	if w == nil {
		t.Error("NewWatchat() should not return nil")
		return
	}
	w.Start()
	time.Sleep(time.Second * 10)
	err = w.AddSource(&config.SourceConfig{
		Name: "mysqlTest",
		Type: "mysql",
		Config: map[string]interface{}{
			"dsn": "root:123456@@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	err = w.AddDriver(&config.DriverConfig{
		Name:       "mysqlTest",
		MetricType: "gauge",
		Labels: map[string]string{
			"group":   "arch",
			"service": "elance",
			"module":  "goods",
		},
		Help: "test help",
		Type: "watchat.driver.mysql",
		Config: map[string]interface{}{
			"query": "select count(*) from users where id > ?",
			"bind_var": []interface{}{
				0,
			},
			"interval": "3s",
			"source":   "mysqlTest",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(time.Second * 10)
	w.Stop()
}

func TestWatchat_Metrics(t *testing.T) {
	var cfg *config.Config
	b, err := ioutil.ReadFile("./testdata/config.yaml")
	if err != nil {
		t.Error(err)
	}
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		t.Error(err)
	}
	w, err := NewWatchat(cfg)
	if err != nil {
		t.Error(err)
		return
	}
	if w == nil {
		t.Error("NewWatchat() should not return nil")
		return
	}
	w.Start()
	time.Sleep(time.Second * 20)

	w.Stop()
}
