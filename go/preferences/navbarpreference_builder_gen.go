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
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("NavbarPreference", err)...)
	}

	if len(errs) != 0 {
		return NavbarPreference{}, errs
	}

	return *builder.internal, nil
}

func (builder *NavbarPreferenceBuilder) SavedItemIds(savedItemIds []string) *NavbarPreferenceBuilder {
	builder.internal.SavedItemIds = savedItemIds

	return builder
}

func (builder *NavbarPreferenceBuilder) applyDefaults() {
}
