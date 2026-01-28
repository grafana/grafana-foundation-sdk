// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package news

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Options] = (*OptionsBuilder)(nil)

type OptionsBuilder struct {
	internal *Options
	errors   cog.BuildErrors
}

func NewOptionsBuilder() *OptionsBuilder {
	resource := NewOptions()
	builder := &OptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *OptionsBuilder) Build() (Options, error) {
	if err := builder.internal.Validate(); err != nil {
		return Options{}, err
	}

	if len(builder.errors) > 0 {
		return Options{}, cog.MakeBuildErrors("news.options", builder.errors)
	}

	return *builder.internal, nil
}

// empty/missing will default to grafana blog
func (builder *OptionsBuilder) FeedUrl(feedUrl string) *OptionsBuilder {
	builder.internal.FeedUrl = &feedUrl

	return builder
}

func (builder *OptionsBuilder) ShowImage(showImage bool) *OptionsBuilder {
	builder.internal.ShowImage = &showImage

	return builder
}
