// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package grafanapyroscope

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[variants.Dataquery] = (*DataqueryBuilder)(nil)

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

func (builder *DataqueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dataquery{}, err
	}

	return *builder.internal, nil
}

// Specifies the query label selectors.
func (builder *DataqueryBuilder) LabelSelector(labelSelector string) *DataqueryBuilder {
	builder.internal.LabelSelector = labelSelector

	return builder
}

// Specifies the query span selectors.
func (builder *DataqueryBuilder) SpanSelector(spanSelector []string) *DataqueryBuilder {
	builder.internal.SpanSelector = spanSelector

	return builder
}

// Specifies the type of profile to query.
func (builder *DataqueryBuilder) ProfileTypeId(profileTypeId string) *DataqueryBuilder {
	builder.internal.ProfileTypeId = profileTypeId

	return builder
}

// Allows to group the results.
func (builder *DataqueryBuilder) GroupBy(groupBy []string) *DataqueryBuilder {
	builder.internal.GroupBy = groupBy

	return builder
}

// Sets the maximum number of nodes in the flamegraph.
func (builder *DataqueryBuilder) MaxNodes(maxNodes int64) *DataqueryBuilder {
	builder.internal.MaxNodes = &maxNodes

	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder {
	builder.internal.RefId = refId

	return builder
}

// true if query is disabled (ie should not be returned to the dashboard)
// Note this does not always imply that the query should not be executed since
// the results from a hidden query may be used as the input to other queries (SSE etc)
func (builder *DataqueryBuilder) Hide(hide bool) *DataqueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *DataqueryBuilder) QueryType(queryType string) *DataqueryBuilder {
	builder.internal.QueryType = &queryType

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

func (builder *DataqueryBuilder) applyDefaults() {
	builder.LabelSelector("{}")
}
