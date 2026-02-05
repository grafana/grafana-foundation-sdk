// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package resource

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Metadata] = (*MetadataBuilder)(nil)

type MetadataBuilder struct {
	internal *Metadata
	errors   cog.BuildErrors
}

func NewMetadataBuilder() *MetadataBuilder {
	resource := NewMetadata()
	builder := &MetadataBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MetadataBuilder) Build() (Metadata, error) {
	if err := builder.internal.Validate(); err != nil {
		return Metadata{}, err
	}

	if len(builder.errors) > 0 {
		return Metadata{}, cog.MakeBuildErrors("resource.metadata", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MetadataBuilder) Name(name string) *MetadataBuilder {
	builder.internal.Name = name

	return builder
}

func (builder *MetadataBuilder) Namespace(namespace string) *MetadataBuilder {
	builder.internal.Namespace = &namespace

	return builder
}

func (builder *MetadataBuilder) Labels(labels map[string]string) *MetadataBuilder {
	builder.internal.Labels = labels

	return builder
}

func (builder *MetadataBuilder) Annotations(annotations map[string]string) *MetadataBuilder {
	builder.internal.Annotations = annotations

	return builder
}

func (builder *MetadataBuilder) Uid(uid string) *MetadataBuilder {
	builder.internal.Uid = &uid

	return builder
}

func (builder *MetadataBuilder) ResourceVersion(resourceVersion string) *MetadataBuilder {
	builder.internal.ResourceVersion = &resourceVersion

	return builder
}

func (builder *MetadataBuilder) Generation(generation int64) *MetadataBuilder {
	builder.internal.Generation = &generation

	return builder
}

func (builder *MetadataBuilder) CreationTimestamp(creationTimestamp string) *MetadataBuilder {
	builder.internal.CreationTimestamp = &creationTimestamp

	return builder
}

func (builder *MetadataBuilder) UpdateTimestamp(updateTimestamp string) *MetadataBuilder {
	builder.internal.UpdateTimestamp = &updateTimestamp

	return builder
}

func (builder *MetadataBuilder) DeletionTimestamp(deletionTimestamp string) *MetadataBuilder {
	builder.internal.DeletionTimestamp = &deletionTimestamp

	return builder
}
