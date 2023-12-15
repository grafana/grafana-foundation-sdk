// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package stat

import (
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

type Options struct {
	GraphMode     common.BigValueGraphMode      `json:"graphMode"`
	ColorMode     common.BigValueColorMode      `json:"colorMode"`
	JustifyMode   common.BigValueJustifyMode    `json:"justifyMode"`
	ReduceOptions common.ReduceDataOptions      `json:"reduceOptions"`
	Text          *common.VizTextDisplayOptions `json:"text,omitempty"`
	TextMode      common.BigValueTextMode       `json:"textMode"`
	Orientation   common.VizOrientation         `json:"orientation"`
}
