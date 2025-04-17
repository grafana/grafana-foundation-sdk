// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package resource

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Manifest] = (*ManifestBuilder)(nil)

type ManifestBuilder struct {
	internal *Manifest
	errors   map[string]cog.BuildErrors
}

func NewManifestBuilder() *ManifestBuilder {
	resource := NewManifest()
	builder := &ManifestBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ManifestBuilder) Build() (Manifest, error) {
	if err := builder.internal.Validate(); err != nil {
		return Manifest{}, err
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
		builder.errors["metadata"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Metadata = metadataResource

	return builder
}

func (builder *ManifestBuilder) Spec(spec any) *ManifestBuilder {
	builder.internal.Spec = spec

	return builder
}
