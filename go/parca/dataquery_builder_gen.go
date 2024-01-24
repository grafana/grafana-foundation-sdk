// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package parca

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

func (builder *DataqueryBuilder) LabelSelector(labelSelector string) *DataqueryBuilder {
	builder.internal.LabelSelector = &labelSelector

	return builder
}

func (builder *DataqueryBuilder) ProfileTypeId(profileTypeId string) *DataqueryBuilder {
	builder.internal.ProfileTypeId = &profileTypeId

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
	builder.LabelSelector("{}")
}
