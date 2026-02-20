// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package preferencesv1alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NavbarPreference] = (*NavbarPreferenceBuilder)(nil)

type NavbarPreferenceBuilder struct {
	internal *NavbarPreference
	errors   cog.BuildErrors
}

func NewNavbarPreferenceBuilder() *NavbarPreferenceBuilder {
	resource := NewNavbarPreference()
	builder := &NavbarPreferenceBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *NavbarPreferenceBuilder) Build() (NavbarPreference, error) {
	if err := builder.internal.Validate(); err != nil {
		return NavbarPreference{}, err
	}

	if len(builder.errors) > 0 {
		return NavbarPreference{}, cog.MakeBuildErrors("preferencesv1alpha1.navbarPreference", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *NavbarPreferenceBuilder) RecordError(path string, err error) *NavbarPreferenceBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *NavbarPreferenceBuilder) BookmarkUrls(bookmarkUrls []string) *NavbarPreferenceBuilder {
	builder.internal.BookmarkUrls = bookmarkUrls

	return builder
}
