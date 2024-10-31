// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferences

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NavbarPreference] = (*NavbarPreferenceBuilder)(nil)

type NavbarPreferenceBuilder struct {
	internal *NavbarPreference
	errors   map[string]cog.BuildErrors
}

func NewNavbarPreferenceBuilder() *NavbarPreferenceBuilder {
	resource := &NavbarPreference{}
	builder := &NavbarPreferenceBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *NavbarPreferenceBuilder) Build() (NavbarPreference, error) {
	if err := builder.internal.Validate(); err != nil {
		return NavbarPreference{}, err
	}

	return *builder.internal, nil
}

func (builder *NavbarPreferenceBuilder) BookmarkUrls(bookmarkUrls []string) *NavbarPreferenceBuilder {
	builder.internal.BookmarkUrls = bookmarkUrls

	return builder
}

func (builder *NavbarPreferenceBuilder) applyDefaults() {
}
