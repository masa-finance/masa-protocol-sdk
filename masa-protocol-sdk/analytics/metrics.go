package analytics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define your metrics by creating Prometheus metrics objects.
var (
	// myMetric is a GaugeVec for tracking a value that can go up and down.
	myMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "my_metric_name",
			Help: "Description of my metric",
		},
		[]string{"label1", "label2"}, // Labels for the metric, if any.
	)
)

func init() {
	// Register the metric with Prometheus's default registry.
	prometheus.MustRegister(myMetric)
}

// TrackMetric updates the gauge for the given metric name with the specified value and labels.
// TrackMetric updates the gauge for the given metric name with the specified value and labels.
func TrackMetric(metricName string, labels map[string]string, value float64) {
	// Convert map of labels to a slice of label values maintaining the order of labels as defined in the metric.
	labelValues := make([]string, 0, len(labels))
	for _, labelName := range []string{"label1", "label2"} {
		labelValues = append(labelValues, labels[labelName])
	}

	// Set the value with the label values directly on the GaugeVec.
	myMetric.WithLabelValues(labelValues...).Set(value)
}

// ReportMetrics starts an HTTP server and exposes the registered metrics at `/metrics` endpoint.
func ReportMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":2112", nil)) // Use your preferred port for Prometheus metrics.
}

// Additional functions and types related to Prometheus metrics can be added here.
// No saves to S3
// Promethus running within AWS - reaches out to nodes to pull info every 10/20 secs
// Defines metrics we are looking for, expose them through an external port
// Consul keeps a record of nodes running so promethus knows who to scrape
// Most important part of this - picking out where to place the trackers and utilizing them right
// Request comes thru for work - someone asks for twitter api endpoint, listening for when that occurs, general health of the nodes / uptime,etc.
// Protocol level info trakced by prometheus, if we scrape twitter data its sent to S3
// Fire event when work begins, understand which scraper, timestamp, workers associated, as work is done - send event when raw data is received, send request when raw data is structured
// this will helpo us understand latency and optimization of work, recereate state at any point in time, recreate network function in any point of time
// Promethus for oracle level health metrics, uptime o fnodes, request metrics
// Setting up promthesus events for network health is main priority
// Class with promethesus and other events
