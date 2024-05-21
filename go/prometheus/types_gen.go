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
	// The actual expression/query that will be evaluated by Prometheus
	Expr *string `json:"expr,omitempty"`
	// Returns only the latest value that Prometheus has scraped for the requested time series
	Instant *bool `json:"instant,omitempty"`
	// Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
	Range *bool `json:"range,omitempty"`
	// Execute an additional query to identify interesting raw samples relevant for the given expr
	Exemplar *bool `json:"exemplar,omitempty"`
	// Specifies which editor is being used to prepare the query. It can be "code" or "builder"
	EditorMode *QueryEditorMode `json:"editorMode,omitempty"`
	// Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
	Format *PromQueryFormat `json:"format,omitempty"`
	// Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
	LegendFormat *string `json:"legendFormat,omitempty"`
	// @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
	// See https://github.com/grafana/grafana/issues/48081
	IntervalFactor *float64 `json:"intervalFactor,omitempty"`
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId *string `json:"refId,omitempty"`
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	Hide *bool `json:"hide,omitempty"`
	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource any `json:"datasource,omitempty"`
	// An additional lower limit for the step parameter of the Prometheus query and for the
	// `$__interval` and `$__rate_interval` variables.
	Interval *string `json:"interval,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}
