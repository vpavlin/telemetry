package prometheus

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/status-im/telemetry/lib/common"
	"github.com/status-im/telemetry/pkg/types"

	"github.com/prometheus/client_golang/prometheus"
	prom_model "github.com/prometheus/client_model/go"
)

type MetricName string

const (
	RegularQueryHistoryMessages MetricName = "waku2_envelopes_validated_total"
)

type MetricWithName struct {
	Name   string
	Metric *prom_model.Metric
}

type PrometheusMetric struct {
	metricsCache map[string]MetricWithName
}

func (r *PrometheusMetric) Process(db *sql.DB, errs *common.MetricErrors, data *types.TelemetryRequest) error {
	d, _ := data.TelemetryData.MarshalJSON()
	fmt.Println(string(d))

	var pm types.PrometheusMetricData
	err := json.Unmarshal(*data.TelemetryData, &pm)
	if err != nil {
		return err
	}

	fmt.Println(pm)

	if r.metricsCache == nil {
		r.metricsCache = make(map[string]MetricWithName)
	}
	for _, v := range pm.Value {
		name := "peerId"
		v.Label = append(v.Label, &prom_model.LabelPair{Name: &name, Value: &pm.PeerID})
		r.metricsCache[hashLabels(pm.Name, v.Label)] = MetricWithName{Name: pm.Name, Metric: v}
	}

	switch MetricName(pm.Name) {
	case RegularQueryHistoryMessages:
		return ProcessReqularMissingMsgQuery(db, errs, &pm)
	}

	return nil
}

func (r *PrometheusMetric) Clean(db *sql.DB, before int64) (int64, error) {
	var resultErr error
	res1, err := common.Cleanup(db, "regularHistoryQuery", before)

	resultErr = errors.Join(err)

	return res1, resultErr
}

func (pm *PrometheusMetric) ServeMetrics(w http.ResponseWriter, r *http.Request) {

	type MetricData struct {
		Name   string
		Value  float64
		Labels []*prom_model.LabelPair
	}

	myTemplate := `{{ range . }}{{ .Name }}{{print "{" }}{{ range .Labels }}{{.Name}}="{{.Value}}" {{ end -}}{{print "}" }} {{ .Value}}
{{end}}`

	data := []MetricData{}

	for _, value := range pm.metricsCache {

		metric := value.Metric
		var v float64

		if metric.Counter != nil {
			v = *metric.Counter.Value
		}

		if metric.Gauge != nil {
			v = *metric.Gauge.Value
		}

		if metric.Summary != nil {
			v = float64(*metric.Summary.SampleCount)
		}

		data = append(data, MetricData{
			Name:   value.Name,
			Labels: metric.Label,
			Value:  v,
		})

	}

	gatherer := prometheus.DefaultGatherer
	metrics, err := gatherer.Gather()
	if err != nil {
		log.Fatalf("Failed to gather metrics: %v", err)
	}

	for _, mf := range metrics {
		metricValue := mf.GetMetric()

		if len(metricValue) == 0 {
			continue
		}

		for _, m := range mf.GetMetric() {
			var v float64
			if m.Counter != nil {
				v = *m.Counter.Value
			}

			if m.Gauge != nil {
				v = *m.Gauge.Value
			}

			if m.Summary != nil {
				v = float64(*m.Summary.SampleCount)
			}

			data = append(data, MetricData{
				Name:   *mf.Name,
				Labels: m.GetLabel(),
				Value:  v,
			})
		}
	}

	tmpl, err := template.New("metrics").Parse(myTemplate)
	if err != nil {
		log.Println("failed to parse template", err)
		return
	}

	tmpl.Execute(w, data)

}

func hashLabels(name string, labels []*prom_model.LabelPair) string {
	toHash := name

	for _, l := range labels {
		toHash += *l.Name + *l.Value
	}

	h := sha256.New()
	h.Write([]byte(toHash))

	return hex.EncodeToString(h.Sum(nil))
}
