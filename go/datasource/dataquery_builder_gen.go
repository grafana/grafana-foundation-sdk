// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package datasource

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[variants.Dataquery] = (*DataqueryBuilder)(nil)

type DataqueryBuilder struct {
	internal *Dataquery
	errors   cog.BuildErrors
}

func NewDataqueryBuilder() *DataqueryBuilder {
	resource := NewDataquery()
	builder := &DataqueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DataqueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dataquery{}, err
	}

	if len(builder.errors) > 0 {
		return Dataquery{}, cog.MakeBuildErrors("datasource.dataquery", builder.errors)
	}

	return *builder.internal, nil
}

// Panel ID from wich the queries will be reused.
func (builder *DataqueryBuilder) PanelId(panelId uint32) *DataqueryBuilder {
	builder.internal.PanelId = panelId

	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder {
	builder.internal.RefId = refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *DataqueryBuilder) Hide(hide bool) *DataqueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *DataqueryBuilder) WithTransforms(withTransforms bool) *DataqueryBuilder {
	builder.internal.WithTransforms = withTransforms

	return builder
}

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *DataqueryBuilder) Datasource(datasource dashboard.DataSourceRef) *DataqueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}
