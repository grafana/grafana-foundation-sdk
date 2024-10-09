// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package grafanapyroscope

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type PyroscopeQueryType string

const (
	PyroscopeQueryTypeMetrics PyroscopeQueryType = "metrics"
	PyroscopeQueryTypeProfile PyroscopeQueryType = "profile"
	PyroscopeQueryTypeBoth    PyroscopeQueryType = "both"
)

type Dataquery struct {
	// Specifies the query label selectors.
	LabelSelector string `json:"labelSelector"`
	// Specifies the query span selectors.
	SpanSelector []string `json:"spanSelector,omitempty"`
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
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	Hide *bool `json:"hide,omitempty"`
	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}

func (resource Dataquery) DataqueryType() string {
	return "grafanapyroscope"
}

func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "grafanapyroscope",
		DataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &Dataquery{}

			if err := json.Unmarshal(raw, dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
		GoConverter: func(input any) string {
			var dataquery Dataquery
			if cast, ok := input.(*Dataquery); ok {
				dataquery = *cast
			} else {
				dataquery = input.(Dataquery)
			}
			return DataqueryConverter(dataquery)
		},
	}
}

func (resource Dataquery) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(Dataquery)
	if !ok {
		return false
	}
	if resource.LabelSelector != other.LabelSelector {
		return false
	}

	if len(resource.SpanSelector) != len(other.SpanSelector) {
		return false
	}

	for i1 := range resource.SpanSelector {
		if resource.SpanSelector[i1] != other.SpanSelector[i1] {
			return false
		}
	}
	if resource.ProfileTypeId != other.ProfileTypeId {
		return false
	}

	if len(resource.GroupBy) != len(other.GroupBy) {
		return false
	}

	for i1 := range resource.GroupBy {
		if resource.GroupBy[i1] != other.GroupBy[i1] {
			return false
		}
	}
	if resource.MaxNodes == nil && other.MaxNodes != nil || resource.MaxNodes != nil && other.MaxNodes == nil {
		return false
	}

	if resource.MaxNodes != nil {
		if *resource.MaxNodes != *other.MaxNodes {
			return false
		}
	}
	if resource.RefId != other.RefId {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}
	if resource.QueryType == nil && other.QueryType != nil || resource.QueryType != nil && other.QueryType == nil {
		return false
	}

	if resource.QueryType != nil {
		if *resource.QueryType != *other.QueryType {
			return false
		}
	}
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
		}
	}

	return true
}
