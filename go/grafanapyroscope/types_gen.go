// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package grafanapyroscope

type PhlareQueryType string

const (
	PhlareQueryTypeMetrics PhlareQueryType = "metrics"
	PhlareQueryTypeProfile PhlareQueryType = "profile"
	PhlareQueryTypeBoth    PhlareQueryType = "both"
)

type Dataquery struct {
	LabelSelector *string  `json:"labelSelector,omitempty"`
	ProfileTypeId *string  `json:"profileTypeId,omitempty"`
	GroupBy       []string `json:"groupBy,omitempty"`
	MaxNodes      *int64   `json:"maxNodes,omitempty"`
	RefId         *string  `json:"refId,omitempty"`
	Hide          *bool    `json:"hide,omitempty"`
	QueryType     *string  `json:"queryType,omitempty"`
	Datasource    any      `json:"datasource,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}
