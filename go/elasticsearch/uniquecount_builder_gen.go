// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[UniqueCount] = (*UniqueCountBuilder)(nil)

type UniqueCountBuilder struct {
	internal *UniqueCount
	errors   map[string]cog.BuildErrors
}

func NewUniqueCountBuilder() *UniqueCountBuilder {
	resource := &UniqueCount{}
	builder := &UniqueCountBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "cardinality"

	return builder
}

func (builder *UniqueCountBuilder) Build() (UniqueCount, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("UniqueCount", err)...)
	}

	if len(errs) != 0 {
		return UniqueCount{}, errs
	}

	return *builder.internal, nil
}

func (builder *UniqueCountBuilder) Field(field string) *UniqueCountBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *UniqueCountBuilder) Id(id string) *UniqueCountBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *UniqueCountBuilder) Settings(settings struct {
	PrecisionThreshold *string `json:"precision_threshold,omitempty"`
	Missing            *string `json:"missing,omitempty"`
}) *UniqueCountBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *UniqueCountBuilder) Hide(hide bool) *UniqueCountBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *UniqueCountBuilder) applyDefaults() {
}
