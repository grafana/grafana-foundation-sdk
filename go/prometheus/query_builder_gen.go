// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

var _ cog.Builder[dashboardv2beta1.DataQueryKind] = (*QueryBuilder)(nil)

type QueryBuilder struct {
	internal *dashboardv2beta1.DataQueryKind
	errors   cog.BuildErrors
}

func NewQueryBuilder() *QueryBuilder {
	resource := dashboardv2beta1.NewDataQueryKind()
	builder := &QueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "DataQuery"
	builder.internal.Group = "prometheus"

	return builder
}

func (builder *QueryBuilder) Build() (dashboardv2beta1.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2beta1.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2beta1.DataQueryKind{}, cog.MakeBuildErrors("prometheus.query", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryBuilder) RecordError(path string, err error) *QueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *QueryBuilder) Version(version string) *QueryBuilder {
	builder.internal.Version = version

	return builder
}

// New type for datasource reference
// Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder {
	datasourceResource, err := datasource.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Datasource = &datasourceResource

	return builder
}

// The actual expression/query that will be evaluated by Prometheus
func (builder *QueryBuilder) Expr(expr string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Expr = expr

	return builder
}

// Returns only the latest value that Prometheus has scraped for the requested time series
func (builder *QueryBuilder) Instant() *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	valInstant := true
	builder.internal.Spec.(*Dataquery).Instant = &valInstant
	valRange := false
	builder.internal.Spec.(*Dataquery).Range = &valRange

	return builder
}

// Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
func (builder *QueryBuilder) Range() *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	valRange := true
	builder.internal.Spec.(*Dataquery).Range = &valRange
	valInstant := false
	builder.internal.Spec.(*Dataquery).Instant = &valInstant

	return builder
}

// Execute an additional query to identify interesting raw samples relevant for the given expr
func (builder *QueryBuilder) Exemplar(exemplar bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Exemplar = &exemplar

	return builder
}

// Specifies which editor is being used to prepare the query. It can be "code" or "builder"
func (builder *QueryBuilder) EditorMode(editorMode QueryEditorMode) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).EditorMode = &editorMode

	return builder
}

// Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
func (builder *QueryBuilder) Format(format PromQueryFormat) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Format = &format

	return builder
}

// Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
func (builder *QueryBuilder) LegendFormat(legendFormat string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).LegendFormat = &legendFormat

	return builder
}

// @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
// See https://github.com/grafana/grafana/issues/48081
func (builder *QueryBuilder) IntervalFactor(intervalFactor float64) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).IntervalFactor = &intervalFactor

	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *QueryBuilder) RefId(refId string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).RefId = &refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *QueryBuilder) Hide(hide bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).QueryType = &queryType

	return builder
}

// An additional lower limit for the step parameter of the Prometheus query and for the
// `$__interval` and `$__rate_interval` variables.
func (builder *QueryBuilder) Interval(interval string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Interval = &interval

	return builder
}

func (builder *QueryBuilder) RangeAndInstant() *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	valRange := true
	builder.internal.Spec.(*Dataquery).Range = &valRange
	valInstant := true
	builder.internal.Spec.(*Dataquery).Instant = &valInstant

	return builder
}
