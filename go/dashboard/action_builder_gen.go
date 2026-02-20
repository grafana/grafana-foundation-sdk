// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Action] = (*ActionBuilder)(nil)

// Dashboard action
type ActionBuilder struct {
	internal *Action
	errors   cog.BuildErrors
}

func NewActionBuilder() *ActionBuilder {
	resource := NewAction()
	builder := &ActionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ActionBuilder) Build() (Action, error) {
	if err := builder.internal.Validate(); err != nil {
		return Action{}, err
	}

	if len(builder.errors) > 0 {
		return Action{}, cog.MakeBuildErrors("dashboard.action", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ActionBuilder) RecordError(path string, err error) *ActionBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ActionBuilder) Type(typeArg ActionType) *ActionBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *ActionBuilder) Title(title string) *ActionBuilder {
	builder.internal.Title = title

	return builder
}

func (builder *ActionBuilder) Fetch(fetch cog.Builder[FetchOptions]) *ActionBuilder {
	fetchResource, err := fetch.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Fetch = &fetchResource

	return builder
}

func (builder *ActionBuilder) Infinity(infinity cog.Builder[InfinityOptions]) *ActionBuilder {
	infinityResource, err := infinity.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Infinity = &infinityResource

	return builder
}

func (builder *ActionBuilder) Confirmation(confirmation string) *ActionBuilder {
	builder.internal.Confirmation = &confirmation

	return builder
}

func (builder *ActionBuilder) OneClick(oneClick bool) *ActionBuilder {
	builder.internal.OneClick = &oneClick

	return builder
}

func (builder *ActionBuilder) Variables(variables []cog.Builder[ActionVariable]) *ActionBuilder {
	variablesResources := make([]ActionVariable, 0, len(variables))
	for _, r1 := range variables {
		variablesDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		variablesResources = append(variablesResources, variablesDepth1)
	}
	builder.internal.Variables = variablesResources

	return builder
}

func (builder *ActionBuilder) Style(style cog.Builder[DashboardActionStyle]) *ActionBuilder {
	styleResource, err := style.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Style = &styleResource

	return builder
}
