// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package debug

type UpdateConfig struct {
	Render        bool `json:"render"`
	DataChanged   bool `json:"dataChanged"`
	SchemaChanged bool `json:"schemaChanged"`
}

type DebugMode string

const (
	DebugModeRender     DebugMode = "render"
	DebugModeEvents     DebugMode = "events"
	DebugModeCursor     DebugMode = "cursor"
	DebugModeState      DebugMode = "State"
	DebugModeThrowError DebugMode = "ThrowError"
)

type Options struct {
	Mode     DebugMode     `json:"mode"`
	Counters *UpdateConfig `json:"counters,omitempty"`
}
