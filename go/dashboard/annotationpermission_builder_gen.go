// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AnnotationPermission] = (*AnnotationPermissionBuilder)(nil)

type AnnotationPermissionBuilder struct {
	internal *AnnotationPermission
	errors   map[string]cog.BuildErrors
}

func NewAnnotationPermissionBuilder() *AnnotationPermissionBuilder {
	resource := NewAnnotationPermission()
	builder := &AnnotationPermissionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *AnnotationPermissionBuilder) Build() (AnnotationPermission, error) {
	if err := builder.internal.Validate(); err != nil {
		return AnnotationPermission{}, err
	}

	return *builder.internal, nil
}

func (builder *AnnotationPermissionBuilder) Dashboard(dashboard cog.Builder[AnnotationActions]) *AnnotationPermissionBuilder {
	dashboardResource, err := dashboard.Build()
	if err != nil {
		builder.errors["dashboard"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Dashboard = &dashboardResource

	return builder
}

func (builder *AnnotationPermissionBuilder) Organization(organization cog.Builder[AnnotationActions]) *AnnotationPermissionBuilder {
	organizationResource, err := organization.Build()
	if err != nil {
		builder.errors["organization"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Organization = &organizationResource

	return builder
}