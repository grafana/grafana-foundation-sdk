// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package prometheus

type QueryEditorMode string

const (
	QueryEditorModeCode    QueryEditorMode = "code"
	QueryEditorModeBuilder QueryEditorMode = "builder"
)

type PromQueryFormat string

const (
	PromQueryFormatTimeSeries PromQueryFormat = "time_series"
	PromQueryFormatTable      PromQueryFormat = "table"
	PromQueryFormatHeatmap    PromQueryFormat = "heatmap"
)

type Dataquery struct {
	Expr           *string          `json:"expr,omitempty"`
	Instant        *bool            `json:"instant,omitempty"`
	Range          *bool            `json:"range,omitempty"`
	Exemplar       *bool            `json:"exemplar,omitempty"`
	EditorMode     *QueryEditorMode `json:"editorMode,omitempty"`
	Format         *PromQueryFormat `json:"format,omitempty"`
	LegendFormat   *string          `json:"legendFormat,omitempty"`
	IntervalFactor *float64         `json:"intervalFactor,omitempty"`
	RefId          *string          `json:"refId,omitempty"`
	Hide           *bool            `json:"hide,omitempty"`
	QueryType      *string          `json:"queryType,omitempty"`
	Datasource     any              `json:"datasource,omitempty"`
	// An additional lower limit for the step parameter of the Prometheus query and for the
	// `$__interval` and `$__rate_interval` variables.
	Interval *string `json:"interval,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}
