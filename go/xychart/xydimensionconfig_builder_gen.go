// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package xychart

import (
	"errors"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[XYDimensionConfig] = (*XYDimensionConfigBuilder)(nil)

// Configuration for the Table/Auto mode
type XYDimensionConfigBuilder struct {
	internal *XYDimensionConfig
	errors   map[string]cog.BuildErrors
}

func NewXYDimensionConfigBuilder() *XYDimensionConfigBuilder {
	resource := &XYDimensionConfig{}
	builder := &XYDimensionConfigBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *XYDimensionConfigBuilder) Build() (XYDimensionConfig, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("XYDimensionConfig", err)...)
	}

	if len(errs) != 0 {
		return XYDimensionConfig{}, errs
	}

	return *builder.internal, nil
}

func (builder *XYDimensionConfigBuilder) Frame(frame int32) *XYDimensionConfigBuilder {
	if !(frame >= 0) {
		builder.errors["frame"] = cog.MakeBuildErrors("frame", errors.New("frame must be >= 0"))
		return builder
	}
	builder.internal.Frame = frame

	return builder
}

func (builder *XYDimensionConfigBuilder) X(x string) *XYDimensionConfigBuilder {
	builder.internal.X = &x

	return builder
}

func (builder *XYDimensionConfigBuilder) Exclude(exclude []string) *XYDimensionConfigBuilder {
	builder.internal.Exclude = exclude

	return builder
}

func (builder *XYDimensionConfigBuilder) applyDefaults() {
}
