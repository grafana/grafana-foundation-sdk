// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package annotationslist

type Options struct {
	OnlyFromThisDashboard bool     `json:"onlyFromThisDashboard"`
	OnlyInTimeRange       bool     `json:"onlyInTimeRange"`
	Tags                  []string `json:"tags"`
	Limit                 uint32   `json:"limit"`
	ShowUser              bool     `json:"showUser"`
	ShowTime              bool     `json:"showTime"`
	ShowTags              bool     `json:"showTags"`
	NavigateToPanel       bool     `json:"navigateToPanel"`
	NavigateBefore        string   `json:"navigateBefore"`
	NavigateAfter         string   `json:"navigateAfter"`
}
