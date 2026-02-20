// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package candlestick

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CandlestickFieldMap] = (*CandlestickFieldMapBuilder)(nil)

type CandlestickFieldMapBuilder struct {
	internal *CandlestickFieldMap
	errors   cog.BuildErrors
}

func NewCandlestickFieldMapBuilder() *CandlestickFieldMapBuilder {
	resource := NewCandlestickFieldMap()
	builder := &CandlestickFieldMapBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *CandlestickFieldMapBuilder) Build() (CandlestickFieldMap, error) {
	if err := builder.internal.Validate(); err != nil {
		return CandlestickFieldMap{}, err
	}

	if len(builder.errors) > 0 {
		return CandlestickFieldMap{}, cog.MakeBuildErrors("candlestick.candlestickFieldMap", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *CandlestickFieldMapBuilder) RecordError(path string, err error) *CandlestickFieldMapBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
