// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardlist

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
		return Options{}, cog.MakeBuildErrors("dashboardlist.options", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsBuilder) KeepTime(keepTime bool) *OptionsBuilder {
	builder.internal.KeepTime = keepTime

	return builder
}

func (builder *OptionsBuilder) IncludeVars(includeVars bool) *OptionsBuilder {
	builder.internal.IncludeVars = includeVars

	return builder
}

func (builder *OptionsBuilder) ShowStarred(showStarred bool) *OptionsBuilder {
	builder.internal.ShowStarred = showStarred

	return builder
}

func (builder *OptionsBuilder) ShowRecentlyViewed(showRecentlyViewed bool) *OptionsBuilder {
	builder.internal.ShowRecentlyViewed = showRecentlyViewed

	return builder
}

func (builder *OptionsBuilder) ShowSearch(showSearch bool) *OptionsBuilder {
	builder.internal.ShowSearch = showSearch

	return builder
}

func (builder *OptionsBuilder) ShowHeadings(showHeadings bool) *OptionsBuilder {
	builder.internal.ShowHeadings = showHeadings

	return builder
}

func (builder *OptionsBuilder) MaxItems(maxItems int64) *OptionsBuilder {
	builder.internal.MaxItems = maxItems

	return builder
}

func (builder *OptionsBuilder) Query(query string) *OptionsBuilder {
	builder.internal.Query = query

	return builder
}

func (builder *OptionsBuilder) Tags(tags []string) *OptionsBuilder {
	builder.internal.Tags = tags

	return builder
}

// folderId is deprecated, and migrated to folderUid on panel init
func (builder *OptionsBuilder) FolderId(folderId int64) *OptionsBuilder {
	builder.internal.FolderId = &folderId

	return builder
}

func (builder *OptionsBuilder) FolderUID(folderUID string) *OptionsBuilder {
	builder.internal.FolderUID = &folderUID

	return builder
}
