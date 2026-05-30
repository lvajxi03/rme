package metrics

import (
	"fmt"
	"strings"

	"github.com/lvajxi03/rme/utils"
)

type MetricType int

const (
	MetricGauge MetricType = iota
	MetricCounter
	// MetricHistogram
	// MetricSummary
	MetricCount
)

type Metric[T int | float64] struct {
	Name       string
	Help       string
	Kind       MetricType
	Value      int
	MinV, MaxV T
	Labels     map[string]string
}

func CreateMetric[T int | float64](name, help string, kind MetricType, values ...T) *Metric[T] {
	m := &Metric[T]{
		Name:   name,
		Help:   help,
		Kind:   kind,
		Value:  0,
		Labels: make(map[string]string),
	}
	if len(values) > 1 {
		m.MinV = values[0]
		m.MaxV = values[1]
	}
	return m
}

func (m *Metric[T]) AddLabel(name, value string) {
	m.Labels[name] = value
}

func (m *Metric[T]) Render() string {

	nazwy_metryk := map[MetricType]string{
		MetricGauge:   "gauge",
		MetricCounter: "counter",
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# HELP %s %s\n# TYPE %s %s\n", m.Name, m.Help, m.Name, nazwy_metryk[m.Kind]))
	sb.WriteString(m.Name)
	sb.WriteString(m.FormatLabels())
	sb.WriteString(" ")
	if m.Kind == MetricCounter {
		m.Value++
		sb.WriteString(fmt.Sprintf("%d\n", m.Value))
	} else {
		sb.WriteString(fmt.Sprintf("%v\n", utils.Wylosuj(m.MinV, m.MaxV)))
	}

	return sb.String()
}

func (m *Metric[T]) FormatLabels() string {
	var sb strings.Builder
	pierwszy := true
	for klucz, wartosc := range m.Labels {
		if !pierwszy {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%s=%s", klucz, wartosc))
		pierwszy = false
	}
	return fmt.Sprintf("{%s}", sb.String())
}

func RenderAllMetrics[T int | float64](all_metrics []*Metric[T]) string {
	var sb strings.Builder
	for _, metric := range all_metrics {
		sb.WriteString(metric.Render())
		sb.WriteString("\n")
	}
	return sb.String()
}
