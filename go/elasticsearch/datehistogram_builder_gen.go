// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DateHistogram] = (*DateHistogramBuilder)(nil)

type DateHistogramBuilder struct {
	internal *DateHistogram
	errors   map[string]cog.BuildErrors
}

func NewDateHistogramBuilder() *DateHistogramBuilder {
	resource := &DateHistogram{}
	builder := &DateHistogramBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "date_histogram"

	return builder
}

func (builder *DateHistogramBuilder) Build() (DateHistogram, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("DateHistogram", err)...)
	}

	if len(errs) != 0 {
		return DateHistogram{}, errs
	}

	return *builder.internal, nil
}

func (builder *DateHistogramBuilder) Field(field string) *DateHistogramBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *DateHistogramBuilder) Id(id string) *DateHistogramBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *DateHistogramBuilder) Settings(settings struct {
	Interval    *string `json:"interval,omitempty"`
	MinDocCount *string `json:"min_doc_count,omitempty"`
	TrimEdges   *string `json:"trimEdges,omitempty"`
	Offset      *string `json:"offset,omitempty"`
	TimeZone    *string `json:"timeZone,omitempty"`
}) *DateHistogramBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *DateHistogramBuilder) applyDefaults() {
}
