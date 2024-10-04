// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package candlestick

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CandlestickColors] = (*CandlestickColorsBuilder)(nil)

type CandlestickColorsBuilder struct {
	internal *CandlestickColors
	errors   map[string]cog.BuildErrors
}

func NewCandlestickColorsBuilder() *CandlestickColorsBuilder {
	resource := &CandlestickColors{}
	builder := &CandlestickColorsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *CandlestickColorsBuilder) Build() (CandlestickColors, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("CandlestickColors", err)...)
	}

	if len(errs) != 0 {
		return CandlestickColors{}, errs
	}

	return *builder.internal, nil
}

func (builder *CandlestickColorsBuilder) Up(up string) *CandlestickColorsBuilder {
	builder.internal.Up = up

	return builder
}

func (builder *CandlestickColorsBuilder) Down(down string) *CandlestickColorsBuilder {
	builder.internal.Down = down

	return builder
}

func (builder *CandlestickColorsBuilder) Flat(flat string) *CandlestickColorsBuilder {
	builder.internal.Flat = flat

	return builder
}

func (builder *CandlestickColorsBuilder) applyDefaults() {
	builder.Up("green")
	builder.Down("red")
	builder.Flat("gray")
}
