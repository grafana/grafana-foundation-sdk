// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

var _ cog.Builder[cogvariants.Dataquery] = (*DataqueryBuilder)(nil)

type DataqueryBuilder struct {
	internal *Dataquery
	errors   map[string]cog.BuildErrors
}

func NewDataqueryBuilder() *DataqueryBuilder {
	resource := &Dataquery{}
	builder := &DataqueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *DataqueryBuilder) Build() (cogvariants.Dataquery, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Dataquery", err)...)
	}

	if len(errs) != 0 {
		return Dataquery{}, errs
	}

	return *builder.internal, nil
}

func (builder *DataqueryBuilder) Expr(expr string) *DataqueryBuilder {
	builder.internal.Expr = &expr

	return builder
}

func (builder *DataqueryBuilder) Instant(instant bool) *DataqueryBuilder {
	builder.internal.Instant = &instant

	return builder
}

func (builder *DataqueryBuilder) Range(rangeArg bool) *DataqueryBuilder {
	builder.internal.Range = &rangeArg

	return builder
}

func (builder *DataqueryBuilder) Exemplar(exemplar bool) *DataqueryBuilder {
	builder.internal.Exemplar = &exemplar

	return builder
}

func (builder *DataqueryBuilder) EditorMode(editorMode QueryEditorMode) *DataqueryBuilder {
	builder.internal.EditorMode = &editorMode

	return builder
}

func (builder *DataqueryBuilder) Format(format PromQueryFormat) *DataqueryBuilder {
	builder.internal.Format = &format

	return builder
}

func (builder *DataqueryBuilder) LegendFormat(legendFormat string) *DataqueryBuilder {
	builder.internal.LegendFormat = &legendFormat

	return builder
}

func (builder *DataqueryBuilder) IntervalFactor(intervalFactor float64) *DataqueryBuilder {
	builder.internal.IntervalFactor = &intervalFactor

	return builder
}

func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder {
	builder.internal.RefId = &refId

	return builder
}

func (builder *DataqueryBuilder) Hide(hide bool) *DataqueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *DataqueryBuilder) QueryType(queryType string) *DataqueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

func (builder *DataqueryBuilder) Datasource(datasource any) *DataqueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

func (builder *DataqueryBuilder) applyDefaults() {
}
