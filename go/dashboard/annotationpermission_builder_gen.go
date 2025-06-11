// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AnnotationPermission] = (*AnnotationPermissionBuilder)(nil)

// +k8s:deepcopy-gen=true
type AnnotationPermissionBuilder struct {
	internal *AnnotationPermission
	errors   cog.BuildErrors
}

func NewAnnotationPermissionBuilder() *AnnotationPermissionBuilder {
	resource := NewAnnotationPermission()
	builder := &AnnotationPermissionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AnnotationPermissionBuilder) Build() (AnnotationPermission, error) {
	if err := builder.internal.Validate(); err != nil {
		return AnnotationPermission{}, err
	}

	if len(builder.errors) > 0 {
		return AnnotationPermission{}, cog.MakeBuildErrors("dashboard.annotationPermission", builder.errors)
	}

	return *builder.internal, nil
}

// +k8s:deepcopy-gen=true
func (builder *AnnotationPermissionBuilder) Dashboard(dashboard cog.Builder[AnnotationActions]) *AnnotationPermissionBuilder {
	dashboardResource, err := dashboard.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Dashboard = &dashboardResource

	return builder
}

// +k8s:deepcopy-gen=true
func (builder *AnnotationPermissionBuilder) Organization(organization cog.Builder[AnnotationActions]) *AnnotationPermissionBuilder {
	organizationResource, err := organization.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Organization = &organizationResource

	return builder
}
