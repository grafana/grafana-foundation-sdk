// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package rolebinding

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RoleBinding] = (*RoleBindingBuilder)(nil)

type RoleBindingBuilder struct {
	internal *RoleBinding
	errors   cog.BuildErrors
}

func NewRoleBindingBuilder() *RoleBindingBuilder {
	resource := NewRoleBinding()
	builder := &RoleBindingBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *RoleBindingBuilder) Build() (RoleBinding, error) {
	if err := builder.internal.Validate(); err != nil {
		return RoleBinding{}, err
	}

	if len(builder.errors) > 0 {
		return RoleBinding{}, cog.MakeBuildErrors("rolebinding.roleBinding", builder.errors)
	}

	return *builder.internal, nil
}

// The role we are discussing
func (builder *RoleBindingBuilder) Role(role BuiltinRoleRefOrCustomRoleRef) *RoleBindingBuilder {
	builder.internal.Role = role

	return builder
}

// The team or user that has the specified role
func (builder *RoleBindingBuilder) Subject(subject cog.Builder[RoleBindingSubject]) *RoleBindingBuilder {
	subjectResource, err := subject.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Subject = subjectResource

	return builder
}
