// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package grafanapyroscope

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
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

// VariantConfig returns the configuration related to grafanapyroscope dataqueries.
// This configuration describes how to unmarshal it, convert it to code, …
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
		StrictDataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &Dataquery{}

			if err := dataquery.UnmarshalJSONStrict(raw); err != nil {
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dataquery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Dataquery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "labelSelector"
	if fields["labelSelector"] != nil {
		if string(fields["labelSelector"]) != "null" {
			if err := json.Unmarshal(fields["labelSelector"], &resource.LabelSelector); err != nil {
				errs = append(errs, cog.MakeBuildErrors("labelSelector", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("labelSelector", errors.New("required field is null"))...)

		}
		delete(fields, "labelSelector")
	} else {
		errs = append(errs, cog.MakeBuildErrors("labelSelector", errors.New("required field is missing from input"))...)
	}
	// Field "spanSelector"
	if fields["spanSelector"] != nil {
		if string(fields["spanSelector"]) != "null" {

			if err := json.Unmarshal(fields["spanSelector"], &resource.SpanSelector); err != nil {
				errs = append(errs, cog.MakeBuildErrors("spanSelector", err)...)
			}

		}
		delete(fields, "spanSelector")

	}
	// Field "profileTypeId"
	if fields["profileTypeId"] != nil {
		if string(fields["profileTypeId"]) != "null" {
			if err := json.Unmarshal(fields["profileTypeId"], &resource.ProfileTypeId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("profileTypeId", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("profileTypeId", errors.New("required field is null"))...)

		}
		delete(fields, "profileTypeId")
	} else {
		errs = append(errs, cog.MakeBuildErrors("profileTypeId", errors.New("required field is missing from input"))...)
	}
	// Field "groupBy"
	if fields["groupBy"] != nil {
		if string(fields["groupBy"]) != "null" {

			if err := json.Unmarshal(fields["groupBy"], &resource.GroupBy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("groupBy", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("groupBy", errors.New("required field is null"))...)

		}
		delete(fields, "groupBy")
	} else {
		errs = append(errs, cog.MakeBuildErrors("groupBy", errors.New("required field is missing from input"))...)
	}
	// Field "maxNodes"
	if fields["maxNodes"] != nil {
		if string(fields["maxNodes"]) != "null" {
			if err := json.Unmarshal(fields["maxNodes"], &resource.MaxNodes); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxNodes", err)...)
			}

		}
		delete(fields, "maxNodes")

	}
	// Field "refId"
	if fields["refId"] != nil {
		if string(fields["refId"]) != "null" {
			if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refId", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is null"))...)

		}
		delete(fields, "refId")
	} else {
		errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is missing from input"))...)
	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}
	// Field "queryType"
	if fields["queryType"] != nil {
		if string(fields["queryType"]) != "null" {
			if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryType", err)...)
			}

		}
		delete(fields, "queryType")

	}
	// Field "datasource"
	if fields["datasource"] != nil {
		if string(fields["datasource"]) != "null" {

			resource.Datasource = &dashboard.DataSourceRef{}
			if err := resource.Datasource.UnmarshalJSONStrict(fields["datasource"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
			}

		}
		delete(fields, "datasource")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Dataquery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
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

// Validate checks all the validation constraints that may be defined on `Dataquery` fields for violations and returns them.
func (resource Dataquery) Validate() error {
	var errs cog.BuildErrors
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}
