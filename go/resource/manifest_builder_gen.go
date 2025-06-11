// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package resource

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Manifest] = (*ManifestBuilder)(nil)

type ManifestBuilder struct {
	internal *Manifest
	errors   cog.BuildErrors
}

func NewManifestBuilder() *ManifestBuilder {
	resource := NewManifest()
	builder := &ManifestBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ManifestBuilder) Build() (Manifest, error) {
	if err := builder.internal.Validate(); err != nil {
		return Manifest{}, err
	}

	if len(builder.errors) > 0 {
		return Manifest{}, cog.MakeBuildErrors("resource.manifest", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ManifestBuilder) ApiVersion(apiVersion string) *ManifestBuilder {
	builder.internal.ApiVersion = apiVersion

	return builder
}

func (builder *ManifestBuilder) Kind(kind string) *ManifestBuilder {
	builder.internal.Kind = kind

	return builder
}

func (builder *ManifestBuilder) Metadata(metadata cog.Builder[Metadata]) *ManifestBuilder {
	metadataResource, err := metadata.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Metadata = metadataResource

	return builder
}

func (builder *ManifestBuilder) Spec(spec any) *ManifestBuilder {
	builder.internal.Spec = spec

	return builder
}
