// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[StreamingQuery] = (*StreamingQueryBuilder)(nil)

type StreamingQueryBuilder struct {
	internal *StreamingQuery
	errors   cog.BuildErrors
}

func NewStreamingQueryBuilder() *StreamingQueryBuilder {
	resource := NewStreamingQuery()
	builder := &StreamingQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *StreamingQueryBuilder) Build() (StreamingQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return StreamingQuery{}, err
	}

	if len(builder.errors) > 0 {
		return StreamingQuery{}, cog.MakeBuildErrors("testdata.streamingQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *StreamingQueryBuilder) Type(typeArg StreamingQueryType) *StreamingQueryBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *StreamingQueryBuilder) Speed(speed int32) *StreamingQueryBuilder {
	builder.internal.Speed = speed

	return builder
}

func (builder *StreamingQueryBuilder) Spread(spread int32) *StreamingQueryBuilder {
	builder.internal.Spread = spread

	return builder
}

func (builder *StreamingQueryBuilder) Noise(noise int32) *StreamingQueryBuilder {
	builder.internal.Noise = noise

	return builder
}

func (builder *StreamingQueryBuilder) Bands(bands int32) *StreamingQueryBuilder {
	builder.internal.Bands = &bands

	return builder
}

func (builder *StreamingQueryBuilder) Url(url string) *StreamingQueryBuilder {
	builder.internal.Url = &url

	return builder
}
