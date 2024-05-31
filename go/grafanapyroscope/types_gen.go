// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package grafanapyroscope

import (
	"encoding/json"

	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

type PhlareQueryType string

const (
	PhlareQueryTypeMetrics PhlareQueryType = "metrics"
	PhlareQueryTypeProfile PhlareQueryType = "profile"
	PhlareQueryTypeBoth    PhlareQueryType = "both"
)

type Dataquery struct {
	// Specifies the query label selectors.
	LabelSelector string `json:"labelSelector"`
	// Specifies the type of profile to query.
	ProfileTypeId string `json:"profileTypeId"`
	// Allows to group the results.
	GroupBy []string `json:"groupBy"`
	// Sets the maximum number of nodes in the flamegraph.
	MaxNodes *int64 `json:"maxNodes,omitempty"`
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource any `json:"datasource,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}

func VariantConfig() cogvariants.DataqueryConfig {
	return cogvariants.DataqueryConfig{
		Identifier: "grafanapyroscope",
		DataqueryUnmarshaler: func(raw []byte) (cogvariants.Dataquery, error) {
			dataquery := Dataquery{}

			if err := json.Unmarshal(raw, &dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
	}
}
