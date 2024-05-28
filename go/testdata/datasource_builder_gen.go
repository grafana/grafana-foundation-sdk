// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Datasource] = (*DatasourceBuilder)(nil)

type DatasourceBuilder struct {
	internal *Datasource
	errors   map[string]cog.BuildErrors
}

func NewDatasourceBuilder() *DatasourceBuilder {
	resource := &Datasource{}
	builder := &DatasourceBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *DatasourceBuilder) Build() (Datasource, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Datasource", err)...)
	}

	if len(errs) != 0 {
		return Datasource{}, errs
	}

	return *builder.internal, nil
}

// The apiserver version
func (builder *DatasourceBuilder) ApiVersion(apiVersion string) *DatasourceBuilder {
	builder.internal.ApiVersion = &apiVersion

	return builder
}

// The datasource plugin type
func (builder *DatasourceBuilder) Type(typeArg string) *DatasourceBuilder {
	builder.internal.Type = typeArg

	return builder
}

// Datasource UID (NOTE: name in k8s)
func (builder *DatasourceBuilder) Uid(uid string) *DatasourceBuilder {
	builder.internal.Uid = &uid

	return builder
}

func (builder *DatasourceBuilder) applyDefaults() {
}
