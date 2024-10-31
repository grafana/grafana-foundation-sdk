// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package candlestick

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CandlestickFieldMap] = (*CandlestickFieldMapBuilder)(nil)

type CandlestickFieldMapBuilder struct {
	internal *CandlestickFieldMap
	errors   map[string]cog.BuildErrors
}

func NewCandlestickFieldMapBuilder() *CandlestickFieldMapBuilder {
	resource := &CandlestickFieldMap{}
	builder := &CandlestickFieldMapBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *CandlestickFieldMapBuilder) Build() (CandlestickFieldMap, error) {
	if err := builder.internal.Validate(); err != nil {
		return CandlestickFieldMap{}, err
	}

	return *builder.internal, nil
}

// Corresponds to the starting value of the given period
func (builder *CandlestickFieldMapBuilder) Open(open string) *CandlestickFieldMapBuilder {
	builder.internal.Open = &open

	return builder
}

// Corresponds to the highest value of the given period
func (builder *CandlestickFieldMapBuilder) High(high string) *CandlestickFieldMapBuilder {
	builder.internal.High = &high

	return builder
}

// Corresponds to the lowest value of the given period
func (builder *CandlestickFieldMapBuilder) Low(low string) *CandlestickFieldMapBuilder {
	builder.internal.Low = &low

	return builder
}

// Corresponds to the final (end) value of the given period
func (builder *CandlestickFieldMapBuilder) Close(close string) *CandlestickFieldMapBuilder {
	builder.internal.Close = &close

	return builder
}

// Corresponds to the sample count in the given period. (e.g. number of trades)
func (builder *CandlestickFieldMapBuilder) Volume(volume string) *CandlestickFieldMapBuilder {
	builder.internal.Volume = &volume

	return builder
}

func (builder *CandlestickFieldMapBuilder) applyDefaults() {
}
