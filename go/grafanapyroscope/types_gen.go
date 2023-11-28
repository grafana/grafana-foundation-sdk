// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package grafanapyroscope

type PyroscopeQueryType string

const (
	PyroscopeQueryTypeMetrics PyroscopeQueryType = "metrics"
	PyroscopeQueryTypeProfile PyroscopeQueryType = "profile"
	PyroscopeQueryTypeBoth    PyroscopeQueryType = "both"
)

type Dataquery struct {
	LabelSelector *string  `json:"labelSelector,omitempty"`
	SpanSelector  []string `json:"spanSelector,omitempty"`
	ProfileTypeId *string  `json:"profileTypeId,omitempty"`
	GroupBy       []string `json:"groupBy,omitempty"`
	MaxNodes      *int64   `json:"maxNodes,omitempty"`
	RefId         *string  `json:"refId,omitempty"`
	Hide          *bool    `json:"hide,omitempty"`
	QueryType     *string  `json:"queryType,omitempty"`
	Datasource    any      `json:"datasource,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}
