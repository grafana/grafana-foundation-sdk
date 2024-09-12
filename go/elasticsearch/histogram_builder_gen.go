// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Histogram] = (*HistogramBuilder)(nil)

type HistogramBuilder struct {
	internal *Histogram
	errors   map[string]cog.BuildErrors
}

func NewHistogramBuilder() *HistogramBuilder {
	resource := &Histogram{}
	builder := &HistogramBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "histogram"

	return builder
}

func (builder *HistogramBuilder) Build() (Histogram, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Histogram", err)...)
	}

	if len(errs) != 0 {
		return Histogram{}, errs
	}

	return *builder.internal, nil
}

func (builder *HistogramBuilder) Field(field string) *HistogramBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *HistogramBuilder) Id(id string) *HistogramBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *HistogramBuilder) Settings(settings struct {
	Interval    *string `json:"interval,omitempty"`
	MinDocCount *string `json:"min_doc_count,omitempty"`
}) *HistogramBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *HistogramBuilder) applyDefaults() {
}
