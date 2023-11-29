// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Count] = (*CountBuilder)(nil)

type CountBuilder struct {
	internal *Count
	errors   map[string]cog.BuildErrors
}

func NewCountBuilder() *CountBuilder {
	resource := &Count{}
	builder := &CountBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "count"

	return builder
}

func (builder *CountBuilder) Build() (Count, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Count", err)...)
	}

	if len(errs) != 0 {
		return Count{}, errs
	}

	return *builder.internal, nil
}

func (builder *CountBuilder) Id(id string) *CountBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *CountBuilder) Hide(hide bool) *CountBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *CountBuilder) applyDefaults() {
}
