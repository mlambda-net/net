package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type Recorder interface {
	Start()
	Stop()
}

type recorder struct {
	namespace string
	subSystem string
	time      prometheus.Histogram
	counter   prometheus.Counter
	gauge     prometheus.Gauge
	start     time.Time
	version   string
	env       string
	app       string
}

func (r *recorder)  createHistogram(name string) {
	buckets := []float64{50, 100, 250, 500, 1000, 1500, 2000, 2500, 3000, 4000, 5000}
	histo := prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace:   r.namespace,
		Subsystem:   r.subSystem,
		Name:        name,
		ConstLabels: prometheus.Labels{"app": r.app, "version": r.version, "env": r.env},
		Buckets:     buckets,
	})

	r.time = histo

	_ = prometheus.Register(histo)

}

func (r * recorder) createCounter(name string)  {
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace:   r.namespace,
		Subsystem: r.subSystem,
		ConstLabels: prometheus.Labels{"app": r.app, "version": r.version, "env": r.env},
		Name:        name,

	} )

	r.counter = counter
	_ = prometheus.Register(counter)
}

func (r * recorder) createGauge(name string)  {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: r.namespace,
		Subsystem: r.subSystem,
		ConstLabels: prometheus.Labels{"app": r.app, "version": r.version, "env": r.env},
		Name:      name,
	})
	r.gauge = gauge
  _ = prometheus.Register(gauge)
}

func (r *recorder) Start() {
	r.start = time.Now()
	r.counter.Inc()
	r.gauge.Inc()
}

func (r *recorder) Stop() {
	duration := time.Since(r.start)
	elapsed := float64(duration.Milliseconds())
	r.time.Observe(elapsed)
	r.gauge.Dec()
}

func NewRecorder(config *Configuration, name string) Recorder {
	r := &recorder{namespace: config.Metric.Namespace,
		subSystem: config.Metric.SubSystem,
		version: config.App.Version,
		env : config.App.Env,
		app : config.App.Name,
	}
	r.createCounter(fmt.Sprintf("%s_request", name))
	r.createHistogram(fmt.Sprintf("%s_time", name))
	r.createGauge(fmt.Sprintf("%s_active", name))
	return r
}
