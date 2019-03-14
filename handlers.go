package stats

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// httpInFlightGauge is used for indicating the number of in-flight http requests.
	httpInFlightGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "http_requests_in_flight",
		Help: "A gauge of http requests currently being served.",
	})

	// httpCounter is a counter for total http requests with response code and method as labels.
	httpCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "A counter for http requests.",
		},
		[]string{"code", "method"},
	)

	// httpDuration is a histogram metric for http handlers. Used for tracking durations, apdex
	// and quantiles. The default buckets are intended to be a good starting
	// point for most apps but might need to be tailored for specific endpoints.
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "A histogram of latencies for http requests.",
			Buckets: []float64{0.005, 0.01, 0.025, 0.05, 0.1, .25, .5, 1, 2.5, 5, 10},
		},
		[]string{"handler", "method"},
	)
)

func init() {
	prometheus.MustRegister(httpInFlightGauge, httpCounter, httpDuration)
}

// WrapHTTPHandler wraps an http handler with the default http statistics implementation.
func WrapHTTPHandler(name string, h http.Handler) http.Handler {
	return promhttp.InstrumentHandlerInFlight(httpInFlightGauge,
		promhttp.InstrumentHandlerDuration(httpDuration.MustCurryWith(prometheus.Labels{"handler": name}),
			promhttp.InstrumentHandlerCounter(httpCounter, h),
		),
	)
}
