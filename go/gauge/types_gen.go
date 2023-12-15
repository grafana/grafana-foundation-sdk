// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package gauge

import (
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	ShowThresholdLabels  bool                          `json:"showThresholdLabels"`
	ReduceOptions        common.ReduceDataOptions      `json:"reduceOptions"`
	Text                 *common.VizTextDisplayOptions `json:"text,omitempty"`
	ShowThresholdMarkers bool                          `json:"showThresholdMarkers"`
	Orientation          common.VizOrientation         `json:"orientation"`
}
