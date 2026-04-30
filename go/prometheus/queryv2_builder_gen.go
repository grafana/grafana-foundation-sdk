// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2"
)

var _ cog.Builder[dashboardv2.DataQueryKind] = (*QueryV2Builder)(nil)

type QueryV2Builder struct {
	internal *dashboardv2.DataQueryKind
	errors   cog.BuildErrors
}

func NewQueryV2Builder() *QueryV2Builder {
	resource := dashboardv2.NewDataQueryKind()
	builder := &QueryV2Builder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "DataQuery"
	builder.internal.Group = "prometheus"

	return builder
}

func (builder *QueryV2Builder) Build() (dashboardv2.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2.DataQueryKind{}, cog.MakeBuildErrors("prometheus.queryV2", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryV2Builder) RecordError(path string, err error) *QueryV2Builder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *QueryV2Builder) Version(version string) *QueryV2Builder {
	builder.internal.Version = version

	return builder
}

func (builder *QueryV2Builder) Labels(labels map[string]string) *QueryV2Builder {
	builder.internal.Labels = labels

	return builder
}

// New type for datasource reference
// Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
func (builder *QueryV2Builder) Datasource(datasource cog.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) *QueryV2Builder {
	datasourceResource, err := datasource.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Datasource = &datasourceResource

	return builder
}

// The actual expression/query that will be evaluated by Prometheus
func (builder *QueryV2Builder) Expr(expr string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Expr = expr

	return builder
}

// Returns only the latest value that Prometheus has scraped for the requested time series
func (builder *QueryV2Builder) Instant(instant bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Instant = &instant

	return builder
}

// Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
func (builder *QueryV2Builder) Range(rangeArg bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Range = &rangeArg

	return builder
}

// Execute an additional query to identify interesting raw samples relevant for the given expr
func (builder *QueryV2Builder) Exemplar(exemplar bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Exemplar = &exemplar

	return builder
}

// Specifies which editor is being used to prepare the query. It can be "code" or "builder"
func (builder *QueryV2Builder) EditorMode(editorMode QueryEditorMode) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).EditorMode = &editorMode

	return builder
}

// Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
func (builder *QueryV2Builder) Format(format PromQueryFormat) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Format = &format

	return builder
}

// Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
func (builder *QueryV2Builder) LegendFormat(legendFormat string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).LegendFormat = &legendFormat

	return builder
}

// @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
// See https://github.com/grafana/grafana/issues/48081
func (builder *QueryV2Builder) IntervalFactor(intervalFactor float64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).IntervalFactor = &intervalFactor

	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *QueryV2Builder) RefId(refId string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).RefId = &refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *QueryV2Builder) Hide(hide bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *QueryV2Builder) QueryType(queryType string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).QueryType = &queryType

	return builder
}

// An additional lower limit for the step parameter of the Prometheus query and for the
// `$__interval` and `$__rate_interval` variables.
func (builder *QueryV2Builder) Interval(interval string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Interval = &interval

	return builder
}
