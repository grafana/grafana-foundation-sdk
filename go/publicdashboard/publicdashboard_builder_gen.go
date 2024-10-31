// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package publicdashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PublicDashboard] = (*PublicDashboardBuilder)(nil)

type PublicDashboardBuilder struct {
	internal *PublicDashboard
	errors   map[string]cog.BuildErrors
}

func NewPublicDashboardBuilder() *PublicDashboardBuilder {
	resource := &PublicDashboard{}
	builder := &PublicDashboardBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *PublicDashboardBuilder) Build() (PublicDashboard, error) {
	if err := builder.internal.Validate(); err != nil {
		return PublicDashboard{}, err
	}

	return *builder.internal, nil
}

// Unique public dashboard identifier
func (builder *PublicDashboardBuilder) Uid(uid string) *PublicDashboardBuilder {
	builder.internal.Uid = uid

	return builder
}

// Dashboard unique identifier referenced by this public dashboard
func (builder *PublicDashboardBuilder) DashboardUid(dashboardUid string) *PublicDashboardBuilder {
	builder.internal.DashboardUid = dashboardUid

	return builder
}

// Unique public access token
func (builder *PublicDashboardBuilder) AccessToken(accessToken string) *PublicDashboardBuilder {
	builder.internal.AccessToken = &accessToken

	return builder
}

// Flag that indicates if the public dashboard is enabled
func (builder *PublicDashboardBuilder) IsEnabled(isEnabled bool) *PublicDashboardBuilder {
	builder.internal.IsEnabled = isEnabled

	return builder
}

// Flag that indicates if annotations are enabled
func (builder *PublicDashboardBuilder) AnnotationsEnabled(annotationsEnabled bool) *PublicDashboardBuilder {
	builder.internal.AnnotationsEnabled = annotationsEnabled

	return builder
}

// Flag that indicates if the time range picker is enabled
func (builder *PublicDashboardBuilder) TimeSelectionEnabled(timeSelectionEnabled bool) *PublicDashboardBuilder {
	builder.internal.TimeSelectionEnabled = timeSelectionEnabled

	return builder
}

func (builder *PublicDashboardBuilder) applyDefaults() {
}
