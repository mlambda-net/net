package metrics

import (
	"net/http"
)

type Metric interface {
	Trace(next http.Handler) http.Handler
	AddMetric(path string, name string)
}

type metric struct {
	recorders map[string]Recorder
	config    *Configuration
}

func (m *metric) AddMetric(path string, name string) {
	m.recorders[path] = NewRecorder(m.config, name)
}

func NewMetric(config *Configuration) Metric {
	return &metric{
		config: config,
		recorders: make(map[string]Recorder),
	}

}

func (m *metric) Trace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rc := m.recorders[r.RequestURI]
		if rc != nil {
			rc.Start()
			next.ServeHTTP(w, r)
			rc.Stop()
		} else {
			next.ServeHTTP(w, r)
		}
	})
}