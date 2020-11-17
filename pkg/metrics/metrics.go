package metrics

import "net/http"

type Metric interface {
	Trace(next http.Handler) http.Handler
}

type metric struct {
	recorders map[string]Recorder
}

func NewMetric(config Configuration, maps map[string]string) Metric {
	m := &metric{recorders: make(map[string]Recorder)}
	for k, v := range maps {
		recorder := NewRecorder(config, v)
		m.recorders[k] = recorder
	}
	return m
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