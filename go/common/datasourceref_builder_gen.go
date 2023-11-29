// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DataSourceRef] = (*DataSourceRefBuilder)(nil)

type DataSourceRefBuilder struct {
	internal *DataSourceRef
	errors   map[string]cog.BuildErrors
}

func NewDataSourceRefBuilder() *DataSourceRefBuilder {
	resource := &DataSourceRef{}
	builder := &DataSourceRefBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *DataSourceRefBuilder) Build() (DataSourceRef, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("DataSourceRef", err)...)
	}

	if len(errs) != 0 {
		return DataSourceRef{}, errs
	}

	return *builder.internal, nil
}

// The plugin type-id
func (builder *DataSourceRefBuilder) Type(typeArg string) *DataSourceRefBuilder {
	builder.internal.Type = &typeArg

	return builder
}

// Specific datasource instance
func (builder *DataSourceRefBuilder) Uid(uid string) *DataSourceRefBuilder {
	builder.internal.Uid = &uid

	return builder
}

func (builder *DataSourceRefBuilder) applyDefaults() {
}
