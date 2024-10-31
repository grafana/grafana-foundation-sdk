// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package rolebinding

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RoleBinding] = (*RoleBindingBuilder)(nil)

type RoleBindingBuilder struct {
	internal *RoleBinding
	errors   map[string]cog.BuildErrors
}

func NewRoleBindingBuilder() *RoleBindingBuilder {
	resource := &RoleBinding{}
	builder := &RoleBindingBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *RoleBindingBuilder) Build() (RoleBinding, error) {
	if err := builder.internal.Validate(); err != nil {
		return RoleBinding{}, err
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
		builder.errors["subject"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Subject = subjectResource

	return builder
}

func (builder *RoleBindingBuilder) applyDefaults() {
}
