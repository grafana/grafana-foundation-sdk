// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package parca

type ParcaQueryType string

const (
	ParcaQueryTypeMetrics ParcaQueryType = "metrics"
	ParcaQueryTypeProfile ParcaQueryType = "profile"
	ParcaQueryTypeBoth    ParcaQueryType = "both"
)

type Dataquery struct {
	LabelSelector *string `json:"labelSelector,omitempty"`
	ProfileTypeId *string `json:"profileTypeId,omitempty"`
	RefId         *string `json:"refId,omitempty"`
	Hide          *bool   `json:"hide,omitempty"`
	QueryType     *string `json:"queryType,omitempty"`
	Datasource    any     `json:"datasource,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}
