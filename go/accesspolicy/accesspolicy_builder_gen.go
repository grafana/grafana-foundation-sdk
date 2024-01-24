// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package accesspolicy

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AccessPolicy] = (*AccessPolicyBuilder)(nil)

type AccessPolicyBuilder struct {
	internal *AccessPolicy
	errors   map[string]cog.BuildErrors
}

func NewAccessPolicyBuilder() *AccessPolicyBuilder {
	resource := &AccessPolicy{}
	builder := &AccessPolicyBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *AccessPolicyBuilder) Build() (AccessPolicy, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("AccessPolicy", err)...)
	}

	if len(errs) != 0 {
		return AccessPolicy{}, errs
	}

	return *builder.internal, nil
}

// The scope where these policies should apply
func (builder *AccessPolicyBuilder) Scope(scope cog.Builder[ResourceRef]) *AccessPolicyBuilder {
	scopeResource, err := scope.Build()
	if err != nil {
		builder.errors["scope"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Scope = scopeResource

	return builder
}

// The role that must apply this policy
func (builder *AccessPolicyBuilder) Role(role cog.Builder[RoleRef]) *AccessPolicyBuilder {
	roleResource, err := role.Build()
	if err != nil {
		builder.errors["role"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Role = roleResource

	return builder
}

// The set of rules to apply.  Note that * is required to modify
// access policy rules, and that "none" will reject all actions
func (builder *AccessPolicyBuilder) Rules(rules cog.Builder[AccessRule]) *AccessPolicyBuilder {
	rulesResource, err := rules.Build()
	if err != nil {
		builder.errors["rules"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Rules = append(builder.internal.Rules, rulesResource)

	return builder
}

func (builder *AccessPolicyBuilder) applyDefaults() {
}
