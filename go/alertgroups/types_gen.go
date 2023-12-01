// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertgroups

type Options struct {
	// Comma-separated list of values used to filter alert results
	Labels string `json:"labels"`
	// Name of the alertmanager used as a source for alerts
	Alertmanager string `json:"alertmanager"`
	// Expand all alert groups by default
	ExpandAll bool `json:"expandAll"`
}
