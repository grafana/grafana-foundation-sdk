// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package rolebinding

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RoleBindingSubject] = (*RoleBindingSubjectBuilder)(nil)

type RoleBindingSubjectBuilder struct {
	internal *RoleBindingSubject
	errors   cog.BuildErrors
}

func NewRoleBindingSubjectBuilder() *RoleBindingSubjectBuilder {
	resource := NewRoleBindingSubject()
	builder := &RoleBindingSubjectBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *RoleBindingSubjectBuilder) Build() (RoleBindingSubject, error) {
	if err := builder.internal.Validate(); err != nil {
		return RoleBindingSubject{}, err
	}

	if len(builder.errors) > 0 {
		return RoleBindingSubject{}, cog.MakeBuildErrors("rolebinding.roleBindingSubject", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RoleBindingSubjectBuilder) Kind(kind RoleBindingSubjectKind) *RoleBindingSubjectBuilder {
	builder.internal.Kind = kind

	return builder
}

// The team/user identifier name
func (builder *RoleBindingSubjectBuilder) Name(name string) *RoleBindingSubjectBuilder {
	builder.internal.Name = name

	return builder
}
