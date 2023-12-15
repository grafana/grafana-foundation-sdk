// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package loki

type QueryEditorMode string

const (
	QueryEditorModeCode    QueryEditorMode = "code"
	QueryEditorModeBuilder QueryEditorMode = "builder"
)

type LokiQueryType string

const (
	LokiQueryTypeRange   LokiQueryType = "range"
	LokiQueryTypeInstant LokiQueryType = "instant"
	LokiQueryTypeStream  LokiQueryType = "stream"
)

type SupportingQueryType string

const (
	SupportingQueryTypeLogsVolume SupportingQueryType = "logsVolume"
	SupportingQueryTypeLogsSample SupportingQueryType = "logsSample"
	SupportingQueryTypeDataSample SupportingQueryType = "dataSample"
)

type LokiQueryDirection string

const (
	LokiQueryDirectionForward  LokiQueryDirection = "forward"
	LokiQueryDirectionBackward LokiQueryDirection = "backward"
)

type Dataquery struct {
	Expr         *string          `json:"expr,omitempty"`
	LegendFormat *string          `json:"legendFormat,omitempty"`
	MaxLines     *int64           `json:"maxLines,omitempty"`
	Resolution   *int64           `json:"resolution,omitempty"`
	EditorMode   *QueryEditorMode `json:"editorMode,omitempty"`
	Range        *bool            `json:"range,omitempty"`
	Instant      *bool            `json:"instant,omitempty"`
	Step         *string          `json:"step,omitempty"`
	RefId        *string          `json:"refId,omitempty"`
	Hide         *bool            `json:"hide,omitempty"`
	QueryType    *string          `json:"queryType,omitempty"`
	Datasource   any              `json:"datasource,omitempty"`
}

func (resource Dataquery) ImplementsDataqueryVariant() {}
