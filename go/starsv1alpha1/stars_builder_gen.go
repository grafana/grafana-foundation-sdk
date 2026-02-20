// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package starsv1alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	resource "github.com/grafana/grafana-foundation-sdk/go/resource"
)

var _ cog.Builder[Stars] = (*StarsBuilder)(nil)

type StarsBuilder struct {
	internal *Stars
	errors   cog.BuildErrors
}

func NewStarsBuilder() *StarsBuilder {
	resource := NewStars()
	builder := &StarsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

// Creates a resource manifest from Stars.
func Manifest(name string, stars cog.Builder[Stars]) *resource.ManifestBuilder {
	builder := resource.NewManifestBuilder()

	builder.ApiVersion("preferences.grafana.app/v1alpha1")

	builder.Metadata(resource.Named(name))

	builder.Kind("Stars")
	starsResource, err := stars.Build()
	if err != nil {
		builder.RecordError("Manifest(stars)", err)
		return builder
	}
	builder.Spec(starsResource)

	return builder
}

func (builder *StarsBuilder) Build() (Stars, error) {
	if err := builder.internal.Validate(); err != nil {
		return Stars{}, err
	}

	if len(builder.errors) > 0 {
		return Stars{}, cog.MakeBuildErrors("starsv1alpha1.stars", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *StarsBuilder) RecordError(path string, err error) *StarsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *StarsBuilder) Resources(resource []cog.Builder[Resource]) *StarsBuilder {
	resourceResources := make([]Resource, 0, len(resource))
	for _, r1 := range resource {
		resourceDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		resourceResources = append(resourceResources, resourceDepth1)
	}
	builder.internal.Resource = resourceResources

	return builder
}

func (builder *StarsBuilder) Resource(resource cog.Builder[Resource]) *StarsBuilder {
	resourceResource, err := resource.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Resource = append(builder.internal.Resource, resourceResource)

	return builder
}
