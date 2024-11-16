// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[StreamingQuery] = (*StreamingQueryBuilder)(nil)

type StreamingQueryBuilder struct {
	internal *StreamingQuery
	errors   map[string]cog.BuildErrors
}

func NewStreamingQueryBuilder() *StreamingQueryBuilder {
	resource := NewStreamingQuery()
	builder := &StreamingQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *StreamingQueryBuilder) Build() (StreamingQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return StreamingQuery{}, err
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
